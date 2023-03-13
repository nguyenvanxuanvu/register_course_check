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
	
	if studentStatus == 2{      // not permit to register student
		return &dto.CheckResponseDTO{
			Status : FAIL,
			StudentStatus: NOT_PERMIT_REGISTER_COURSE,
		}, nil
	}

	if studentStatus != 1{
		return nil, errors.New(common.NOT_FOUND_STUDENT_STATUS)
	}

	// Check course (HT, TQ, SH)
	var courseRegisterList []string
	num_credits := 0
	for _, course := range req.RegisterSubjects{
		if slices.Contains(courseRegisterList, course.SubjectId){
			return nil, errors.New(common.DUPLICATE_COURSE_REGISTER)
		} else {
			courseRegisterList = append(courseRegisterList, course.SubjectId)
			if s.dbConfig.GetSubjectConfig(course.SubjectId) == nil{
				return nil, errors.New(common.NOT_FOUND_SUBJECT_ID + ": " + course.SubjectId)
			}
			num_credits += s.dbConfig.GetSubjectConfig(course.SubjectId).NumCredits
		}
	}

	var subjectCheckResults []*dto.SubjectCheck
	var courseNeedChecks []string

	for _, courseId := range courseRegisterList{
		if s.dbConfig.GetSubjectConfig(courseId).SubjectConditionConfig != nil{
			courseNeedChecks = append(courseNeedChecks, courseId)
		}
	} 
	
	if len(courseNeedChecks) == 0{
		return nil, errors.New(common.NOT_FOUND_COURSE_REGISTER)
	}

	
	for _, courseId := range courseNeedChecks{
		condition := s.dbConfig.GetSubjectConfig(courseId).SubjectConditionConfig
		
		listStudyResult := s.client.GetStudyResult(int(req.StudentId))
		
		subjectCheckResult := &dto.SubjectCheck{
			SubjectId: courseId,
			SubjectName: s.dbConfig.GetSubjectConfig(courseId).SubjectName,
			CheckResult: PASS,
		}
		if CheckConditionRecursion(courseId, subjectCheckResult, &condition.Condition, listStudyResult, courseNeedChecks){
			subjectCheckResults = nil
		} else {
			if subjectCheckResult.CheckResult != "PASS"{
				subjectCheckResults = append(subjectCheckResults, subjectCheckResult)
			}
		}
		
		

	}


	
	// check min credit

	checkMinCreditResult := PASS
	minCreditsConfig, maxCreditsConfig := s.repository.GetMinMaxCredit(req.AcademicProgram, int(req.Semester))
	if (minCreditsConfig <= 0 || maxCreditsConfig <= 0){
		return nil, errors.New(common.NOT_FOUND_MIN_MAX_CREDIT_CONFIG)
	}
	
	if num_credits < minCreditsConfig || num_credits > maxCreditsConfig{
		checkMinCreditResult = FAIL
	}


	status := PASS
	if !(len(subjectCheckResults) == 0 && checkMinCreditResult == PASS){
		status = FAIL
	}

	return &dto.CheckResponseDTO{
		Status: status,
		StudentStatus: NORMAL_STATUS_STUDENT,
		SubjectChecks: subjectCheckResults,
		CheckMinCreditResult: checkMinCreditResult,

	}, nil
}




func CheckConditionRecursion(courseId string, subjectCheckResult *dto.SubjectCheck, c *dto.SubjectCondition, results []client.CourseResult, courseNeedChecks []string) bool{
	if c.Left == nil && c.Right == nil {
		if CheckCondition(courseId, subjectCheckResult, c.Data, results, courseNeedChecks){
			return true
		} else {
			return false
		}
	} else {
		if c.Data == "AND"{
			return CheckConditionRecursion(courseId, subjectCheckResult, c.Left, results, courseNeedChecks) && 
			CheckConditionRecursion(courseId, subjectCheckResult, c.Right, results, courseNeedChecks)
		}else{
			return CheckConditionRecursion(courseId, subjectCheckResult, c.Left, results, courseNeedChecks) ||
			CheckConditionRecursion(courseId, subjectCheckResult, c.Right, results, courseNeedChecks)
		}

	}
}


func CheckCondition(courseId string, subjectCheckResult *dto.SubjectCheck, data string, results []client.CourseResult, courseNeedChecks []string) bool{

	

	dt := strings.Split(data, "-")
	if dt[1] == "1" {
		if !CheckContain(results, dt[0]) || notSuccess(results, dt[0], 1){
			subjectCheckResult.CheckResult = FAIL
			subjectCheckResult.FailReasons = append(subjectCheckResult.FailReasons, &dto.Reason{
				SubjectDesId: dt[0],
				ConditionType: 1,
			})
			return false
		}
		return true
	} else if dt[1] == "2" {
		if !CheckContain(results, dt[0]) || notSuccess(results, dt[0], 2){
			subjectCheckResult.CheckResult = FAIL
			subjectCheckResult.FailReasons = append(subjectCheckResult.FailReasons, &dto.Reason{
				SubjectDesId: dt[0],
				ConditionType: 2,
			})
			return false
		}
		return true
	} else {
		if !slices.Contains(courseNeedChecks, dt[0]){
			subjectCheckResult.CheckResult = FAIL
			subjectCheckResult.FailReasons = append(subjectCheckResult.FailReasons, &dto.Reason{
				SubjectDesId: dt[0],
				ConditionType: 3,
			})
			return false
		} 
		return true
	}

}



func CheckContain(courseResults []client.CourseResult, subjectId string) bool {
	for _, courseResult := range courseResults{
		if courseResult.CourseId == subjectId {
			return true
		}
	}
	return false
}


// theType is the condition check (1,2)
func notSuccess(courseResults []client.CourseResult, courseId string, theType int) bool{
	if theType == 1{
		for _, result := range courseResults {
			if result.CourseId == courseId && result.Result != 1 {
				return true
			}
		}
		return false

	} else {
		for _, result := range courseResults {
			if result.CourseId == courseId && result.Result != 1 && result.Result != 2{
				return true
			}
		}
		return false
	}
	
	
}

