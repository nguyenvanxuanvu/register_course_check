package service

import (
	"context"
	"errors"

	"register_course_check/pkg/common"
	"register_course_check/pkg/dto"
)

const NOT_PERMIT_REGISTER_COURSE = "NOT_PERMIT_REGISTER_COURSE"
const NORMAL_STATUS_STUDENT = "NORMAL"
const FAIL = "FAIL"
const PASS = "PASS"

func (s *registerCourseCheckServiceImp) Check(ctx context.Context, req *dto.CheckRequestDTO) (*dto.CheckResponseDTO, error) {
	
	// check student status 
	studentStatus := s.CheckStudentStatus(int(req.StudentId))
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
		courseRegisterList = append(courseRegisterList, course.SubjectId)
		num_credits += s.dbConfig.GetSubjectConfig(course.SubjectId).NumCredits
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
					listDoneCourse := s.GetListDoneCourse(int(req.StudentId))
					if !CheckContain(listDoneCourse, condition.SubjectDesId){
						subjectCheckResult.CheckResult = FAIL
						subjectCheckResult.FailReasons = append(subjectCheckResult.FailReasons, &dto.Reason{
							SubjectDesId: condition.SubjectDesId,
							ConditionType: condition.ConditionType,
						})
						flag = true

					}
					

				}
				if condition.ConditionType == 3{   //SH
					if !CheckContain(courseNeedChecks, condition.SubjectDesId){
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
	minCreditsConfig := s.repository.GetMinCredit(req.AcademicProgram, int(req.Semester))
	if (minCreditsConfig <= 0){
		return nil, errors.New(common.NOT_FOUND_MIN_CREDIT_CONFIG)
	}
	if num_credits < minCreditsConfig {
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


func (s *registerCourseCheckServiceImp) CheckStudentStatus (studentId int) int {
	return s.repository.GetStudentStatus(studentId)
}

func CheckContain(courseNeedCheck []string, subjectId string) bool {
	for _, courseId := range courseNeedCheck{
		if courseId == subjectId {
			return true
		}
	}
	return false
}

func (s *registerCourseCheckServiceImp) GetListDoneCourse(studentId int) []string{
	return s.repository.GetListDoneCourse(studentId)
}

