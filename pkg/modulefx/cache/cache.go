package cache

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"register_course_check/pkg/modulefx/client"
	"register_course_check/redis"

	"github.com/spf13/viper"

	"go.uber.org/fx"
)


type CacheService interface {
	GetStudyResult(ctx context.Context, studentId int) ([]client.CourseResult, error)
	TrySetStudyResult(ctx context.Context, studentId int, studyResult []client.CourseResult) (bool, error)
	GetStudentStatus(ctx context.Context, studentId int) (int, error)
	TrySetStudentStatus(ctx context.Context, studentId int, studentStatus int) (bool, error)
}


type cacheService struct {
	rdb redis.Cache
}

func NewCacheService(rdb redis.Cache) CacheService {
	return &cacheService{rdb: rdb}
}

var Module = fx.Provide(NewCacheService)


func (c cacheService) GetStudyResult(ctx context.Context, studentId int) ([]client.CourseResult, error) {
	var err error = nil
	

	cacheKey := GetStudyResultCacheKey(studentId)
	studentInfo, err := c.rdb.Get(ctx, cacheKey).Result()
	if studentInfo == "" || err != nil {
		return nil, err
	}
	studyResultModel := []client.CourseResult{}

	err = json.Unmarshal([]byte(studentInfo), &studyResultModel)
	if err != nil {
		return nil, err
	}

	

	return studyResultModel, nil
}

func (c cacheService) TrySetStudyResult(ctx context.Context, studentId int, studyResult []client.CourseResult) (bool, error) {
	var err error = nil

	cacheKey := GetStudyResultCacheKey(studentId)
	ttl := viper.GetInt("student-info-ttl-ms")
	
	success, err := c.rdb.SetNX(ctx, cacheKey, studyResult, time.Millisecond*time.Duration(ttl)).Result()
	if err != nil {
		return false, err
	}

	return success, nil
}

func (c cacheService) GetStudentStatus(ctx context.Context, studentId int) (int, error) {
	var err error = nil
	

	cacheKey := GetStudentStatusCacheKey(studentId)
	status, err := c.rdb.Get(ctx, cacheKey).Result()
	if status == "" || err != nil {
		return -1, err
	}
	studentStatus, err := strconv.Atoi(status)
	if err != nil {
		return -1, err
	}

	

	

	return studentStatus, nil
}

func (c cacheService) TrySetStudentStatus(ctx context.Context, studentId int, studentStatus int) (bool, error) {
	var err error = nil

	cacheKey := GetStudentStatusCacheKey(studentId)
	ttl := viper.GetInt("student-info-ttl-ms")

	success, err := c.rdb.SetNX(ctx, cacheKey, studentStatus , time.Millisecond*time.Duration(ttl)).Result()
	if err != nil {
		return false, err
	}

	return success, nil
}

