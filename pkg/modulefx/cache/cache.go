package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/nguyenvanxuanvu/register_course_check/pkg/modulefx/client"
	"github.com/nguyenvanxuanvu/register_course_check/redis/redisconfig"

	"github.com/spf13/viper"
)


type CacheService interface {
	GetStudyResult(ctx context.Context, studentId int) ([]client.CourseResult, error)
	TrySetStudyResult(ctx context.Context, studentId int, studyResult []client.CourseResult) (bool, error)
	GetStudentInfo(ctx context.Context, studentId int) (*client.StudentInfo, error)
	TrySetStudentInfo(ctx context.Context, studentId int, studentInfo *client.StudentInfo) (bool, error)
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
	studyResult, err := c.rdb.Get(ctx, cacheKey).Result()
	if studyResult == "" || err != nil {
		return nil, err
	}
	studyResultModel := []client.CourseResult{}

	err = json.Unmarshal([]byte(studyResult), &studyResultModel)
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

func (c cacheService) GetStudentInfo(ctx context.Context, studentId int) (*client.StudentInfo, error) {
	var err error = nil
	

	cacheKey := GetStudentInfoCacheKey(studentId)
	studentInfo, err := c.rdb.Get(ctx, cacheKey).Result()
	if studentInfo == "" || err != nil {
		return nil, err
	}
	studentInfoModel := &client.StudentInfo{}
	err = json.Unmarshal([]byte(studentInfo), &studentInfoModel)
	if err != nil {
		return nil, err
	}
	

	return studentInfoModel, nil
}

func (c cacheService) TrySetStudentInfo(ctx context.Context, studentId int, studentInfo *client.StudentInfo) (bool, error) {
	var err error = nil

	cacheKey := GetStudentInfoCacheKey(studentId)
	ttl := viper.GetInt("student-info-ttl-ms")
	data, err := json.Marshal(studentInfo)
	if err != nil {
		return false, nil
	}
	success, err := c.rdb.SetNX(ctx, cacheKey, data , time.Millisecond*time.Duration(ttl)).Result()
	if err != nil {
		return false, err
	}

	return success, nil

}

