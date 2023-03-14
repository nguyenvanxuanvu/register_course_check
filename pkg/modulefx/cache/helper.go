package cache

import (
	"fmt"

	"github.com/spf13/viper"
)

const DELIMITER = ":"

const (
	STUDY_RESULT_CACHE_KEY_PREFIX string = "study_result"
	STUDENT_STATUS_CACHE_KEY_PREFIX string = "student_status"
)

func getCachePrefix() string {
	return viper.GetString("cache.prefix")
}

func GetStudyResultCacheKey(studentId int) string {
	return getCachePrefix() + DELIMITER + STUDY_RESULT_CACHE_KEY_PREFIX + DELIMITER + fmt.Sprint(studentId)
}

func GetStudentStatusCacheKey(studentId int) string {
	return getCachePrefix() + DELIMITER + STUDENT_STATUS_CACHE_KEY_PREFIX + DELIMITER + fmt.Sprint(studentId)
}
