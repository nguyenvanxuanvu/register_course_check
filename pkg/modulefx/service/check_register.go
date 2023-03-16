package service

import (
	"context"
	"errors"
	"github.com/nguyenvanxuanvu/register_course_check/pkg/common"
	"github.com/nguyenvanxuanvu/register_course_check/pkg/dto"
	"github.com/nguyenvanxuanvu/register_course_check/pkg/modulefx/client"
	"strings"

	"golang.org/x/exp/slices"
)

const NOT_PERMIT_REGISTER_COURSE = "NOT_PERMIT_REGISTER_COURSE"
const NORMAL_STATUS_STUDENT = "NORMAL"
const FAIL = "FAIL"
const PASS = "PASS"
const AND  = "AND"
const TQ = "1"
const HT = "2"

func (s *registerCourseCheckServiceImp) Check(ctx context.Context, req *dto.CheckRequestDTO) (*dto.CheckResponseDTO, error) {

	// check student status
	studentId := int(req.StudentId)
	studentStatus, _ := s.cacheService.GetStudentStatus(ctx, studentId)
	if studentStatus == -1 {
		studentStatus = s.client.GetStudentStatus(studentId)
		_ , err := s.cacheService.TrySetStudentStatus(ctx, studentId, studentStatus)
		if err != nil {
			return nil, errors.New(common.SET_STUDENT_STATUS_FAIL_REDIS)
		}
	}
	

	if studentStatus == 2 { // not permit to register student
		return &dto.CheckResponseDTO{
			Status:        FAIL,
			StudentStatus: NOT_PERMIT_REGISTER_COURSE,
		}, nil
	}

	if studentStatus != 1 {
		return nil, errors.New(common.NOT_FOUND_STUDENT_STATUS)
	}

	var courseRegisterList []string
	num_credits := 0
	for _, course := range req.RegisterCourses {
		if slices.Contains(courseRegisterList, course.CourseId) {
			return nil, errors.New(common.DUPLICATE_COURSE_REGISTER)
		} else {
			courseRegisterList = append(courseRegisterList, course.CourseId)
			if s.dbConfig.GetCourseConfig(course.CourseId) == nil {
				return nil, errors.New(common.NOT_FOUND_COURSE_ID + ": " + course.CourseId)
			}
			num_credits += s.dbConfig.GetCourseConfig(course.CourseId).NumCredits
		}
	}

	if len(courseRegisterList) == 0 {
		return nil, errors.New(common.NOT_FOUND_COURSE_REGISTER)
	}

	var courseCheckResults []*dto.CourseCheck
	var courseNeedChecks []string

	for _, courseId := range courseRegisterList {
		if s.dbConfig.GetCourseConfig(courseId).CourseConditionConfig != nil {
			courseNeedChecks = append(courseNeedChecks, courseId)
		}
	}

	

	if len(courseNeedChecks) > 0{
		var listStudyResult []client.CourseResult
		listStudyResult, _ = s.cacheService.GetStudyResult(ctx, studentId)
		if listStudyResult == nil {
			listStudyResult = s.client.GetStudyResult(int(req.StudentId))
			_ , err := s.cacheService.TrySetStudyResult(ctx, studentId, listStudyResult)
			if err != nil {
				return nil, errors.New(common.SET_STUDY_RESULT_FAIL_REDIS)
			}
		}
		

		for _, courseId := range courseNeedChecks {
			condition := s.dbConfig.GetCourseConfig(courseId).CourseConditionConfig


			courseCheckResult := &dto.CourseCheck{
				CourseId:   courseId,
				CourseName: s.dbConfig.GetCourseConfig(courseId).CourseName,
				CheckResult: PASS,
			}
			if !s.CheckConditionRecursion(courseId, courseCheckResult.CourseName, courseCheckResult, &condition.Condition, listStudyResult, courseRegisterList) {
				if courseCheckResult.CheckResult != PASS {
					courseCheckResults = append(courseCheckResults, courseCheckResult)
				}
			}

		}
	}

	

	

	// check min credit

	checkMinCreditResult := PASS
	checkMaxCreditResult := PASS
	minCreditsConfig, maxCreditsConfig := s.repository.GetMinMaxCredit(req.AcademicProgram, int(req.Semester))

	if minCreditsConfig < 0 {
		return nil, errors.New(common.NOT_FOUND_MIN_CREDIT_CONFIG)
	}
	if maxCreditsConfig < 0 {
		return nil, errors.New(common.NOT_FOUND_MAX_CREDIT_CONFIG)
	}

	if num_credits < minCreditsConfig {
		checkMinCreditResult = FAIL
	}

	if num_credits > maxCreditsConfig {
		checkMaxCreditResult = FAIL
	}

	if minCreditsConfig > maxCreditsConfig {
		return nil, errors.New(common.MIN_MAX_CONFIG_WRONG)
	}

	status := PASS
	if !(len(courseCheckResults) == 0 && checkMinCreditResult == PASS && checkMaxCreditResult == PASS) {
		status = FAIL
	}

	return &dto.CheckResponseDTO{
		Status:               status,
		StudentStatus:        NORMAL_STATUS_STUDENT,
		CourseChecks:         courseCheckResults,
		CheckMinCreditResult: checkMinCreditResult,
		CheckMaxCreditResult: checkMaxCreditResult,
	}, nil
}





// recursion of check condition
// return true if pass condition

func (s *registerCourseCheckServiceImp) CheckConditionRecursion(courseId string, courseName string, courseCheckResult *dto.CourseCheck, c *dto.CourseCondition, results []client.CourseResult, courseRegisterList []string) bool {
	if c.Left == nil && c.Right == nil {
		if s.CheckCondition(courseId, courseName, courseCheckResult, c.Data, results, courseRegisterList) {
			return true
		} else {
			return false
		}
	} else {
		leftResult := s.CheckConditionRecursion(courseId, courseName, courseCheckResult, c.Left, results, courseRegisterList)
		rightResult := s.CheckConditionRecursion(courseId, courseName, courseCheckResult, c.Right, results, courseRegisterList)
		if c.Data == AND {

			return leftResult && rightResult

		} else {

			return leftResult || rightResult
		}

	}
}

// check condition of each object (data)
// return true if pass condition
func (s *registerCourseCheckServiceImp) CheckCondition(courseId string, courseName string, courseCheckResult *dto.CourseCheck, data string, results []client.CourseResult, courseRegisterList []string) bool {

	dt := strings.Split(data, "-")
	if dt[1] == TQ {

		if !CheckContain(results, dt[0]) || !isSuccess(results, dt[0], 1) {
			courseCheckResult.CheckResult = FAIL
			courseCheckResult.FailReasons = append(courseCheckResult.FailReasons, &dto.Reason{
				CourseDesId:   dt[0],
				CourseDesName: s.dbConfig.GetCourseConfig(dt[0]).CourseName,
				ConditionType:  1,
			})

			return false
		}
		return true
	} else if dt[1] == HT {

		if !CheckContain(results, dt[0]) || !isSuccess(results, dt[0], 2) {
			courseCheckResult.CheckResult = FAIL
			courseCheckResult.FailReasons = append(courseCheckResult.FailReasons, &dto.Reason{
				CourseDesId:   dt[0],
				CourseDesName: s.dbConfig.GetCourseConfig(dt[0]).CourseName,
				ConditionType:  2,
			})

			return false
		}
		return true
	} else {
		if !slices.Contains(courseRegisterList, dt[0]) && !isSuccess(results, dt[0], 2) {
			courseCheckResult.CheckResult = FAIL
			courseCheckResult.FailReasons = append(courseCheckResult.FailReasons, &dto.Reason{
				CourseDesId:   dt[0],
				CourseDesName: s.dbConfig.GetCourseConfig(dt[0]).CourseName,
				ConditionType:  3,
			})
			return false
		}
		return true
	}

}

// check list course result contains a course or not

func CheckContain(courseResults []client.CourseResult, courseId string) bool {
	for _, courseResult := range courseResults {
		if courseResult.CourseId == courseId {
			return true
		}
	}
	return false
}

// theType is the condition check (1,2)
// check a result of course is success or not
// return true if success

func isSuccess(courseResults []client.CourseResult, courseId string, theType int) bool {
	if theType == 1 {
		for _, result := range courseResults {
			if result.CourseId == courseId && (result.Result == 1 || result.Result == 2){
				return true
			}
		}
		return false

	} else {
		for _, result := range courseResults {
			if result.CourseId == courseId {
				return true
			}
		}
		return false
	}

}
