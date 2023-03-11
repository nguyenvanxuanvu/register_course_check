package service

import (
	"context"
	"errors"
	"register_course_check/pkg/common"
	"register_course_check/pkg/modulefx/client"
	"register_course_check/pkg/dto"
	"golang.org/x/exp/slices"
)

const NOT_PERMIT_REGISTER_COURSE = "NOT_PERMIT_REGISTER_COURSE"
const NORMAL_STATUS_STUDENT = "NORMAL"
const FAIL = "FAIL"
const PASS = "PASS"

func (s *registerCourseCheckServiceImp) Check(ctx context.Context, req *dto.CheckRequestDTO) (*dto.CheckResponseDTO, error) {
	
	// check student status 
	studentStatus := s.client.GetStudentStatus(int(req.StudentId))
	if studentStatus == 0{
		return nil, errors.New(common.NOT_FOUND_STUDENT_STATUS)
	}
	if studentStatus == 2{      // not permit to register student
		return &dto.CheckResponseDTO{
			Status : FAIL,
			StudentStatus: NOT_PERMIT_REGISTER_COURSE,
		}, nil
	}

	// Check course (HT, TQ, SH)
	var courseRegisterList []string
	num_credits := 0
	for _, course := range req.RegisterSubjects{
		if !slices.Contains(courseRegisterList, course.SubjectId){
			courseRegisterList = append(courseRegisterList, course.SubjectId)
			if s.dbConfig.GetSubjectConfig(course.SubjectId) == nil{
				return nil, errors.New(common.NOT_FOUND_SUBJECT_ID)
			}
			num_credits += s.dbConfig.GetSubjectConfig(course.SubjectId).NumCredits
		}
		
	}

	var subjectCheckResults []*dto.SubjectCheck
	var courseNeedChecks []string

	for _, courseId := range courseRegisterList{
		if len(s.dbConfig.GetSubjectConfig(courseId).SubjectConditionConfig) > 0{
			courseNeedChecks = append(courseNeedChecks, courseId)
		}
	}

	if len(courseNeedChecks) > 0{
		for _, courseId := range courseNeedChecks{
			
			flag := false
			subjectCheckResult := dto.SubjectCheck{
				SubjectId: courseId,
				SubjectName: s.dbConfig.GetSubjectConfig(courseId).SubjectName,
				CheckResult: PASS,
			}
			for _,condition := range s.dbConfig.GetSubjectConfig(courseId).SubjectConditionConfig{
				if condition.ConditionType == 1 || condition.ConditionType == 2{    // TQ, HT
					listStudyResult := s.client.GetStudyResult(int(req.StudentId))
					if !CheckContain(listStudyResult, condition.SubjectDesId){
						subjectCheckResult.CheckResult = FAIL
						subjectCheckResult.FailReasons = append(subjectCheckResult.FailReasons, &dto.Reason{
							SubjectDesId: condition.SubjectDesId,
							ConditionType: condition.ConditionType,
						})
						flag = true

					}
					

				}
				if condition.ConditionType == 3{   //SH
					if !slices.Contains(courseNeedChecks, condition.SubjectDesId){
						subjectCheckResult.CheckResult = FAIL
						subjectCheckResult.FailReasons = append(subjectCheckResult.FailReasons, &dto.Reason{
							SubjectDesId: condition.SubjectDesId,
							ConditionType: condition.ConditionType,
						})
						flag = true
					}
				}
				
			}
			if flag{
				subjectCheckResults = append(subjectCheckResults, &subjectCheckResult)
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


func CheckContain(courseResults []client.CourseResult, subjectId string) bool {
	for _, courseResult := range courseResults{
		if courseResult.CourseId == subjectId {
			return true
		}
	}
	return false
}


