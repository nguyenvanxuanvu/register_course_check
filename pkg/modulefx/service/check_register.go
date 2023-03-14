package service

import (
	"context"
	"errors"
	"register_course_check/pkg/common"
	"register_course_check/pkg/dto"
	"register_course_check/pkg/modulefx/client"
	"strings"

	"golang.org/x/exp/slices"
)

const NOT_PERMIT_REGISTER_COURSE = "NOT_PERMIT_REGISTER_COURSE"
const NORMAL_STATUS_STUDENT = "NORMAL"
const FAIL = "FAIL"
const PASS = "PASS"

func (s *registerCourseCheckServiceImp) Check(ctx context.Context, req *dto.CheckRequestDTO) (*dto.CheckResponseDTO, error) {

	// check student status
	studentId := int(req.StudentId)
	studentStatus := s.client.GetStudentStatus(studentId)

	if studentStatus == 2 { // not permit to register student
		return &dto.CheckResponseDTO{
			Status:        FAIL,
			StudentStatus: NOT_PERMIT_REGISTER_COURSE,
		}, nil
	}

	if studentStatus != 1 {
		return nil, errors.New(common.NOT_FOUND_STUDENT_STATUS)
	}

	// Check course (HT, TQ, SH)
	var courseRegisterList []string
	num_credits := 0
	for _, course := range req.RegisterSubjects {
		if slices.Contains(courseRegisterList, course.SubjectId) {
			return nil, errors.New(common.DUPLICATE_COURSE_REGISTER)
		} else {
			courseRegisterList = append(courseRegisterList, course.SubjectId)
			if s.dbConfig.GetSubjectConfig(course.SubjectId) == nil {
				return nil, errors.New(common.NOT_FOUND_SUBJECT_ID + ": " + course.SubjectId)
			}
			num_credits += s.dbConfig.GetSubjectConfig(course.SubjectId).NumCredits
		}
	}

	if len(courseRegisterList) == 0 {
		return nil, errors.New(common.NOT_FOUND_COURSE_REGISTER)
	}

	var subjectCheckResults []*dto.SubjectCheck
	var courseNeedChecks []string

	for _, courseId := range courseRegisterList {
		if s.dbConfig.GetSubjectConfig(courseId).SubjectConditionConfig != nil {
			courseNeedChecks = append(courseNeedChecks, courseId)
		}
	}

	for _, courseId := range courseNeedChecks {
		condition := s.dbConfig.GetSubjectConfig(courseId).SubjectConditionConfig

		listStudyResult := s.client.GetStudyResult(int(req.StudentId))

		subjectCheckResult := &dto.SubjectCheck{
			SubjectId:   courseId,
			SubjectName: s.dbConfig.GetSubjectConfig(courseId).SubjectName,
			CheckResult: PASS,
		}
		if !s.CheckConditionRecursion(courseId, subjectCheckResult.SubjectName, subjectCheckResult, &condition.Condition, listStudyResult, courseRegisterList) {
			if subjectCheckResult.CheckResult != PASS {
				subjectCheckResults = append(subjectCheckResults, subjectCheckResult)
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
	if !(len(subjectCheckResults) == 0 && checkMinCreditResult == PASS && checkMaxCreditResult == PASS) {
		status = FAIL
	}

	return &dto.CheckResponseDTO{
		Status:               status,
		StudentStatus:        NORMAL_STATUS_STUDENT,
		SubjectChecks:        subjectCheckResults,
		CheckMinCreditResult: checkMinCreditResult,
		CheckMaxCreditResult: checkMaxCreditResult,
	}, nil
}

// recursion of check condition
// return true if pass condition

func (s *registerCourseCheckServiceImp) CheckConditionRecursion(courseId string, courseName string, subjectCheckResult *dto.SubjectCheck, c *dto.SubjectCondition, results []client.CourseResult, courseRegisterList []string) bool {
	if c.Left == nil && c.Right == nil {
		if s.CheckCondition(courseId, courseName, subjectCheckResult, c.Data, results, courseRegisterList) {
			return true
		} else {
			return false
		}
	} else {
		leftResult := s.CheckConditionRecursion(courseId, courseName, subjectCheckResult, c.Left, results, courseRegisterList)
		rightResult := s.CheckConditionRecursion(courseId, courseName, subjectCheckResult, c.Right, results, courseRegisterList)
		if c.Data == "AND" {

			return leftResult && rightResult

		} else {

			return leftResult || rightResult
		}

	}
}

// check condition of each object (data)
// return true if pass condition
func (s *registerCourseCheckServiceImp) CheckCondition(courseId string, courseName string, subjectCheckResult *dto.SubjectCheck, data string, results []client.CourseResult, courseRegisterList []string) bool {

	dt := strings.Split(data, "-")
	if dt[1] == "1" {

		if !CheckContain(results, dt[0]) || !isSuccess(results, dt[0], 1) {
			subjectCheckResult.CheckResult = FAIL
			subjectCheckResult.FailReasons = append(subjectCheckResult.FailReasons, &dto.Reason{
				SubjectDesId:   dt[0],
				SubjectDesName: s.dbConfig.GetSubjectConfig(dt[0]).SubjectName,
				ConditionType:  1,
			})

			return false
		}
		return true
	} else if dt[1] == "2" {

		if !CheckContain(results, dt[0]) || !isSuccess(results, dt[0], 2) {
			subjectCheckResult.CheckResult = FAIL
			subjectCheckResult.FailReasons = append(subjectCheckResult.FailReasons, &dto.Reason{
				SubjectDesId:   dt[0],
				SubjectDesName: s.dbConfig.GetSubjectConfig(dt[0]).SubjectName,
				ConditionType:  2,
			})

			return false
		}
		return true
	} else {
		if !slices.Contains(courseRegisterList, dt[0]) && !isSuccess(results, dt[0], 2) {
			subjectCheckResult.CheckResult = FAIL
			subjectCheckResult.FailReasons = append(subjectCheckResult.FailReasons, &dto.Reason{
				SubjectDesId:   dt[0],
				SubjectDesName: s.dbConfig.GetSubjectConfig(dt[0]).SubjectName,
				ConditionType:  3,
			})
			return false
		}
		return true
	}

}

// check list course result contains a subject or not

func CheckContain(courseResults []client.CourseResult, subjectId string) bool {
	for _, courseResult := range courseResults {
		if courseResult.CourseId == subjectId {
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
