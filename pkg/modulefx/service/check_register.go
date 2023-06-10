package service

import (
	"context"
	"errors"

	"strconv"

	"github.com/nguyenvanxuanvu/register_course_check/pkg/common"
	"github.com/nguyenvanxuanvu/register_course_check/pkg/dto"
	"github.com/nguyenvanxuanvu/register_course_check/pkg/modulefx/client"

	"golang.org/x/exp/slices"
)

const NOT_PERMIT_REGISTER_COURSE = "NOT_PERMIT_REGISTER_COURSE"
const NORMAL_STATUS_STUDENT = "NORMAL"
const FAIL = "FAIL"
const PASS = "PASS"
const AND = "AND"
const TQ = 1
const HT = 2
const SH = 3

const NORMAL_STATUS = 1
const NOT_PERMIT_REGISTER_STUDENT = 2

func (s *registerCourseCheckServiceImp) Check(ctx context.Context, req *dto.CheckRequestDTO) (*dto.CheckResponseDTO, error) {
	// check student status
	studentId := req.StudentId
	var studentInfo *client.StudentInfo
	studentInfo, err := s.checkStudentStatus(ctx, studentId)
	if err != nil {
		if err.Error() == NOT_PERMIT_REGISTER_COURSE {
			return &dto.CheckResponseDTO{
				Status:        FAIL,
				StudentStatus: NOT_PERMIT_REGISTER_COURSE,
			}, nil
		} else {
			return nil, err
		}
	}

	var courseRegisterList []string
	courseIdToCourseNum := map[string]int{}
	num_credits := 0
	for _, course := range req.RegisterCourses {
		if slices.Contains(courseRegisterList, course.CourseId) {
			return nil, errors.New(common.DUPLICATE_COURSE_REGISTER + ": " + course.CourseId)
		} else {
			courseRegisterList = append(courseRegisterList, course.CourseId)
			courseIdToCourseNum[course.CourseId] = course.CourseNum
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

	// check condition of list course

	listCourseConditionChan := make(chan chanResult[bool], 1)

	go s.checkListCourseConditionAsync(ctx, &courseCheckResults, courseNeedChecks, studentId, courseIdToCourseNum, courseRegisterList, listCourseConditionChan)

	// check min credit

	checkMinCreditResult := dto.MinMaxCredit{
		CheckResult: PASS,
	}
	checkMaxCreditResult := dto.MinMaxCredit{
		CheckResult: PASS,
	}

	minMaxCreditChan := make(chan chanResult[bool], 1)

	go s.checkMinMaxCreditAsync(ctx, &checkMinCreditResult, &checkMaxCreditResult, studentId, studentInfo.AcademicProgram, int(req.Semester), num_credits, minMaxCreditChan)

	minMaxCreditRes, listCourseConditionRes := <-minMaxCreditChan, <-listCourseConditionChan
	if minMaxCreditRes.err != nil || listCourseConditionRes.err != nil {
		return nil, oneOf(minMaxCreditRes.err, listCourseConditionRes.err)
	}

	status := PASS
	if !(len(courseCheckResults) == 0 && checkMinCreditResult.CheckResult == PASS && checkMaxCreditResult.CheckResult == PASS) {
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
	if c.Op == "" {
		if s.CheckCondition(courseId, courseName, courseCheckResult, c.Course, results, courseRegisterList) {
			return true
		} else {
			return false
		}
	} else {
		if c.Op == AND {
			var listCheck []bool
			for _, obj := range c.Leaves {
				checkResult := s.CheckConditionRecursion(courseId, courseName, courseCheckResult, obj, results, courseRegisterList)
				listCheck = append(listCheck, checkResult)
			}

			var returnBool bool = true
			for _, each := range listCheck {
				returnBool = returnBool && each
			}
			return returnBool
		} else {
			for _, obj := range c.Leaves {
				checkResult := s.CheckConditionRecursion(courseId, courseName, courseCheckResult, obj, results, courseRegisterList)
				if checkResult  {
					return true
				}
			}
			return false
		}

	}
}

// check condition of each object (data)
// return true if pass condition
func (s *registerCourseCheckServiceImp) CheckCondition(courseId string, courseName string, courseCheckResult *dto.CourseCheck, data *dto.CourseConditionInfo, results []client.CourseResult, courseRegisterList []string) bool {

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
		if !slices.Contains(courseRegisterList, data.CourseDesId) && !isSuccess(results, data.CourseDesId, HT) {
			courseCheckResult.CheckResult = FAIL
			courseCheckResult.FailReasons = append(courseCheckResult.FailReasons, &dto.Reason{
				CourseDesId:   data.CourseDesId,
				CourseDesName: s.dbConfig.GetCourseConfig(data.CourseDesId).CourseName,
				ConditionType: SH,
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
	if theType == TQ {
		for _, result := range courseResults {
			if result.CourseId == courseId && (result.Result == 1 || result.Result == 2) {
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

func (s *registerCourseCheckServiceImp) checkMinMaxCreditAsync(ctx context.Context, checkMinCreditResult *dto.MinMaxCredit, checkMaxCreditResult *dto.MinMaxCredit, studentId string, academicProgram string, semester int, numCredits int, c chan<- chanResult[bool]) {
	result := chanResult[bool]{}
	result.result = false
	minCreditsConfig, maxCreditsConfig := -1, -1
	var err error
	minMaxCredit, _ := s.cacheService.GetMinMaxCredit(ctx, studentId+"_"+strconv.Itoa(semester))
	if minMaxCredit == nil {
		minCreditsConfig, maxCreditsConfig, err = s.repository.GetMinMaxCredit(studentId, academicProgram, semester)
		if err != nil {
			result.err = err
			c <- result
			return
		}
		_, err := s.cacheService.TrySetMinMaxCredit(ctx, studentId+"_"+strconv.Itoa(semester), []int{minCreditsConfig, maxCreditsConfig})
		if err != nil {
			result.err = errors.New(common.SET_MIN_MAX_CREDIT_FAIL_REDIS)
			c <- result
			return
		}
	} else {
		minCreditsConfig = minMaxCredit[0]
		maxCreditsConfig = minMaxCredit[1]
	}

	if minCreditsConfig < 0 {
		result.err = errors.New(common.MIN_CREDIT_CONFIG_WRONG)
		c <- result
		return
	}
	if maxCreditsConfig < 0 {
		result.err = errors.New(common.MAX_CREDIT_CONFIG_WRONG)
		c <- result
		return
	}

	if numCredits < minCreditsConfig {
		checkMinCreditResult.CheckResult = FAIL
		checkMinCreditResult.CurrentRegister = numCredits
		checkMinCreditResult.Config = minCreditsConfig
	}

	if numCredits > maxCreditsConfig {
		checkMaxCreditResult.CheckResult = FAIL
		checkMaxCreditResult.CurrentRegister = numCredits
		checkMaxCreditResult.Config = maxCreditsConfig
	}

	if minCreditsConfig > maxCreditsConfig {
		result.err = errors.New(common.MIN_MAX_CONFIG_WRONG)
		c <- result
		return
	}
	result.result = true
	c <- result
}

func (s *registerCourseCheckServiceImp) checkListCourseConditionAsync(ctx context.Context, courseCheckResults *([]*dto.CourseCheck), courseNeedChecks []string, studentId string, courseIdToCourseNum map[string]int, courseRegisterList []string, c chan<- chanResult[bool]) {
	result := chanResult[bool]{}
	result.result = false
	if len(courseNeedChecks) > 0 {
		var listStudyResult []client.CourseResult
		listStudyResult, _ = s.cacheService.GetStudyResult(ctx, studentId)
		if listStudyResult == nil {
			listStudyResult = s.client.GetStudyResult(studentId)
			_, err := s.cacheService.TrySetStudyResult(ctx, studentId, listStudyResult)
			if err != nil {
				result.err = errors.New(common.SET_STUDY_RESULT_FAIL_REDIS)
				c <- result
				return
			}
		}

		for _, courseId := range courseNeedChecks {
			condition := s.dbConfig.GetCourseConfig(courseId).CourseConditionConfig

			courseCheckResult := &dto.CourseCheck{
				CourseId:    courseId,
				CourseNum:   courseIdToCourseNum[courseId],
				CourseName:  s.dbConfig.GetCourseConfig(courseId).CourseName,
				CheckResult: PASS,
			}
			if !s.CheckConditionRecursion(courseId, courseCheckResult.CourseName, courseCheckResult, condition.Condition, listStudyResult, courseRegisterList) {
				if courseCheckResult.CheckResult != PASS {
					*courseCheckResults = append(*courseCheckResults, courseCheckResult)
				}
			}

		}
	}

	result.result = true
	c <- result
}

func oneOf(es ...error) error {
	for _, ele := range es {
		if ele != nil {
			return ele
		}
	}
	return nil
}

func (s *registerCourseCheckServiceImp) checkStudentStatus(ctx context.Context, studentId string) (*client.StudentInfo, error) {

	studentInfo, _ := s.cacheService.GetStudentInfo(ctx, studentId)
	studentStatus := -1
	if studentInfo != nil {
		studentStatus = studentInfo.StudentStatus
	}
	if studentStatus == -1 {

		studentInfo = s.client.GetStudentInfo(studentId)
		if studentInfo == nil {
			return nil, errors.New(common.NOT_FOUND_STUDENT_STATUS)
		}
		studentStatus = studentInfo.StudentStatus
		_, err := s.cacheService.TrySetStudentInfo(ctx, studentId, studentInfo)
		if err != nil {
			return nil, errors.New(common.SET_STUDENT_INFO_FAIL_REDIS)
		}
	}

	if studentStatus == NOT_PERMIT_REGISTER_STUDENT { // not permit to register student
		return nil, errors.New(NOT_PERMIT_REGISTER_COURSE)
	}

	if studentStatus != NORMAL_STATUS {
		return nil, errors.New(common.NOT_FOUND_STUDENT_STATUS)
	}
	return studentInfo, nil
}
