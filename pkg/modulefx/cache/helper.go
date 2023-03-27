package cache

import (
	"fmt"

	"github.com/spf13/viper"
)

const DELIMITER = ":"

const (
	STUDY_RESULT_CACHE_KEY_PREFIX string = "study_result"
	STUDENT_INFO_CACHE_KEY_PREFIX string = "student_info"
)

func getCachePrefix() string {
	return viper.GetString("cache.prefix")
}

func GetStudyResultCacheKey(studentId int) string {
	return getCachePrefix() + DELIMITER + STUDY_RESULT_CACHE_KEY_PREFIX + DELIMITER + fmt.Sprint(studentId)
}

func GetStudentInfoCacheKey(studentId int) string {
	return getCachePrefix() + DELIMITER + STUDENT_INFO_CACHE_KEY_PREFIX + DELIMITER + fmt.Sprint(studentId)
}
