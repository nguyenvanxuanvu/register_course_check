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
	GetStudyResult(ctx context.Context, studentId string) ([]client.CourseResult, error)
	TrySetStudyResult(ctx context.Context, studentId string, studyResult []client.CourseResult) (bool, error)
	GetStudentInfo(ctx context.Context, studentId string) (*client.StudentInfo, error)
	TrySetStudentInfo(ctx context.Context, studentId string, studentInfo *client.StudentInfo) (bool, error)

	GetMinMaxCredit(ctx context.Context, studentIdWithSemester string) ([]int, error)
	TrySetMinMaxCredit(ctx context.Context, studentIdWithSemester string, minMaxCredit []int) (bool, error)
}


type cacheService struct {
	rdb redisconfig.Cache
}

func NewCacheService(rdb redisconfig.Cache) CacheService {
	return &cacheService{rdb: rdb}
}


func (c cacheService) GetStudyResult(ctx context.Context, studentId string) ([]client.CourseResult, error) {
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

func (c cacheService) TrySetStudyResult(ctx context.Context, studentId string, studyResult []client.CourseResult) (bool, error) {
	var err error = nil

	cacheKey := GetStudyResultCacheKey(studentId)
	ttl := viper.GetInt("student-ttl-ms")
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

func (c cacheService) GetStudentInfo(ctx context.Context, studentId string) (*client.StudentInfo, error) {
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

func (c cacheService) TrySetStudentInfo(ctx context.Context, studentId string, studentInfo *client.StudentInfo) (bool, error) {
	var err error = nil

	cacheKey := GetStudentInfoCacheKey(studentId)
	ttl := viper.GetInt("student-ttl-ms")
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


func (c cacheService) GetMinMaxCredit(ctx context.Context, studentIdWithSemester string) ([]int, error) {
	var err error = nil
	

	cacheKey := GetMinMaxCreditKey(studentIdWithSemester)
	minMaxCredit, err := c.rdb.Get(ctx, cacheKey).Result()
	if minMaxCredit == "" || err != nil {
		return nil, err
	}
	var minMaxCreditModel []int

	err = json.Unmarshal([]byte(minMaxCredit), &minMaxCreditModel)
	if err != nil {
		return nil, err
	}

	

	return minMaxCreditModel, nil
}


func (c cacheService) TrySetMinMaxCredit(ctx context.Context, studentIdWithSemester string, minMaxCredit []int) (bool, error) {
	var err error = nil

	cacheKey := GetMinMaxCreditKey(studentIdWithSemester)
	ttl := viper.GetInt("min-max-credit-ttl-ms")
	if err != nil {
		return false, nil
	}
	data, err := json.Marshal(minMaxCredit)
	if err != nil {
		return false, nil
	}
	success, err := c.rdb.SetNX(ctx, cacheKey, data , time.Millisecond*time.Duration(ttl)).Result()
	if err != nil {
		return false, err
	}

	return success, nil

}
