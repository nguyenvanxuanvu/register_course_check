package service

import (
	"context"

	"github.com/nguyenvanxuanvu/register_course_check/pkg/dto"
	"github.com/nguyenvanxuanvu/register_course_check/pkg/modulefx/client"
)

func (s *registerCourseCheckServiceImp) checkListCourseConditionForSuggestion(ctx context.Context, courseCheckResults *([]*dto.CourseCheck), courseNeedChecks []string, studentId string, listStudyResult []client.CourseResult) error{
	if len(courseNeedChecks) > 0 {
		for _, courseId := range courseNeedChecks {
			condition := s.dbConfig.GetCourseConfig(courseId).CourseConditionConfig

			courseCheckResult := &dto.CourseCheck{
				CourseId:    courseId,
				CourseName:  s.dbConfig.GetCourseConfig(courseId).CourseName,
				CheckResult: PASS,
			}
			if !s.CheckConditionRecursionForSuggestion(courseId, courseCheckResult.CourseName, courseCheckResult, condition.Condition, listStudyResult) {
				if courseCheckResult.CheckResult != PASS {
					*courseCheckResults = append(*courseCheckResults, courseCheckResult)
				}
			}

		}
	}
	return nil
}

func (s *registerCourseCheckServiceImp) CheckConditionRecursionForSuggestion(courseId string, courseName string, courseCheckResult *dto.CourseCheck, c *dto.CourseCondition, results []client.CourseResult) bool {
	if c.Op == "" {
		if s.CheckConditionForSuggsetion(courseId, courseName, courseCheckResult, c.Course, results) {
			return true
		} else {
			return false
		}
	} else {
		if c.Op == AND {
			var listCheck []bool
			for _, obj := range c.Leaves {
				checkResult := s.CheckConditionRecursionForSuggestion(courseId, courseName, courseCheckResult, obj, results)
				listCheck = append(listCheck, checkResult)
			}

			var returnBool bool = true
			for _, each := range listCheck {
				returnBool = returnBool && each
			}
			return returnBool
		} else {
			for _, obj := range c.Leaves {
				checkResult := s.CheckConditionRecursionForSuggestion(courseId, courseName, courseCheckResult, obj, results)
				if checkResult {
					return true
				}
			}
			return false
		}

	}
}

func (s *registerCourseCheckServiceImp) CheckConditionForSuggsetion(courseId string, courseName string, courseCheckResult *dto.CourseCheck, data *dto.CourseConditionInfo, results []client.CourseResult) bool {

	if data.Type == TQ {

		if !CheckContain(results, data.CourseDesId) || !isSuccess(results, data.CourseDesId, TQ) {
			courseCheckResult.CheckResult = FAIL
			courseCheckResult.FailReasons = append(courseCheckResult.FailReasons, &dto.Reason{
				CourseDesId:   data.CourseDesId,
				CourseDesName: s.dbConfig.GetCourseConfig(data.CourseDesId).CourseName,
				ConditionType: TQ,
			})

			return false
		}
		return true
	} else if data.Type == HT {

		if !CheckContain(results, data.CourseDesId) || !isSuccess(results, data.CourseDesId, HT) {
			courseCheckResult.CheckResult = FAIL
			courseCheckResult.FailReasons = append(courseCheckResult.FailReasons, &dto.Reason{
				CourseDesId:   data.CourseDesId,
				CourseDesName: s.dbConfig.GetCourseConfig(data.CourseDesId).CourseName,
				ConditionType: HT,
			})

			return false
		}
		return true
	} else {
		return true
	}

}
