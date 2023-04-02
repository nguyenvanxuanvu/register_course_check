package cache

import (
	"fmt"

	"github.com/spf13/viper"
)

const DELIMITER = ":"

const (
	STUDY_RESULT_CACHE_KEY_PREFIX string = "study_result"
	STUDENT_INFO_CACHE_KEY_PREFIX string = "student_info"
	MIN_MAX_CREDIT_CACHE_KEY_PREFIX string = "min_max_credit"
)

func getCachePrefix() string {
	return viper.GetString("cache.prefix")
}

func GetStudyResultCacheKey(studentId string) string {
	return getCachePrefix() + DELIMITER + STUDY_RESULT_CACHE_KEY_PREFIX + DELIMITER + fmt.Sprint(studentId)
}

func GetStudentInfoCacheKey(studentId string) string {
	return getCachePrefix() + DELIMITER + STUDENT_INFO_CACHE_KEY_PREFIX + DELIMITER + fmt.Sprint(studentId)
}

func GetMinMaxCreditKey(studentIdWithSemester string) string {
	return getCachePrefix() + DELIMITER + MIN_MAX_CREDIT_CACHE_KEY_PREFIX + DELIMITER + fmt.Sprint(studentIdWithSemester)
}
