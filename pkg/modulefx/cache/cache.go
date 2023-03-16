package cache

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/nguyenvanxuanvu/register_course_check/pkg/modulefx/client"
	"github.com/nguyenvanxuanvu/register_course_check/redis/redisconfig"

	"github.com/spf13/viper"
)


type CacheService interface {
	GetStudyResult(ctx context.Context, studentId int) ([]client.CourseResult, error)
	TrySetStudyResult(ctx context.Context, studentId int, studyResult []client.CourseResult) (bool, error)
	GetStudentStatus(ctx context.Context, studentId int) (int, error)
	TrySetStudentStatus(ctx context.Context, studentId int, studentStatus int) (bool, error)
}


type cacheService struct {
	rdb redisconfig.Cache
}

func NewCacheService(rdb redisconfig.Cache) CacheService {
	return &cacheService{rdb: rdb}
}


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
	data, err := json.Marshal(studyResult)
	if err != nil {
		return false, nil
	}
	success, err := c.rdb.SetNX(ctx, cacheKey, data, time.Millisecond*time.Duration(ttl)).Result()
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

