package service

import (
	"context"
	"errors"
	"strconv"

	"github.com/nguyenvanxuanvu/register_course_check/pkg/common"
	"github.com/nguyenvanxuanvu/register_course_check/pkg/dto"
	"github.com/nguyenvanxuanvu/register_course_check/pkg/modulefx/client"
)

const (
	NOT_PASS      = 1
	TEACHING_PLAN = 2
)

func (s *registerCourseCheckServiceImp) Suggestion(ctx context.Context, req *dto.SuggestionRequestDTO) (*dto.SuggestionResponseDTO, error) {
	studentId := req.StudentId
	var studentInfo *client.StudentInfo
	studentInfo, err := s.checkStudentStatus(ctx, studentId)
	if err != nil {
		return nil, err
	}

	semester := int(req.Semester)
	var failReasons []dto.CourseSuggestion
	var courseSuggestionsFailReasons []dto.CourseSuggestion

	failCoursesChan := make(chan chanResult[bool], 1)
	teachingPlanCoursesChan := make(chan chanResult[bool], 1)
	minMaxCreditChan := make(chan chanResult[bool], 1)

	// get fail course list of student
	go s.getFailCoursesAsync(ctx, studentId, semester, &courseSuggestionsFailReasons, failCoursesChan)

	// get list course of teaching plan of student

	var freeCreditInfo []dto.FreeCreditInfo
	var courseSuggestionsTeachingPlan []dto.CourseSuggestion
	go s.getCoursesTeachingPlan(ctx, studentId, studentInfo, &courseSuggestionsTeachingPlan, &freeCreditInfo, teachingPlanCoursesChan)

	// get min max credit config for student
	var minCredit, maxCredit int
	go s.getMinMaxCreditAsync(ctx, studentId, studentInfo.AcademicProgram, semester, &minCredit, &maxCredit, minMaxCreditChan)

	failCoursesRes, teachingPlanCoursesRes, minMaxCreditRes := <-failCoursesChan, <-teachingPlanCoursesChan, <-minMaxCreditChan
	if failCoursesRes.err != nil || teachingPlanCoursesRes.err != nil || minMaxCreditRes.err != nil {
		return nil, oneOf(failCoursesRes.err, teachingPlanCoursesRes.err, minMaxCreditRes.err)
	}
	failReasons = append(courseSuggestionsFailReasons, courseSuggestionsTeachingPlan...)
	return &dto.SuggestionResponseDTO{
		Courses:          failReasons,
		HintOfFreeCredit: freeCreditInfo,
		MinCredit:        minCredit,
		MaxCredit:        maxCredit,
	}, nil
}

func (s *registerCourseCheckServiceImp) getFailCoursesAsync(ctx context.Context, studentId string, semester int, courseSuggestionsFailReasons *[]dto.CourseSuggestion, c chan<- chanResult[bool]) {
	result := chanResult[bool]{}
	result.result = false

	var listStudyResult []client.CourseResult
	listStudyResult, _ = s.cacheService.GetStudyResult(ctx, studentId+"_"+strconv.Itoa(semester))
	if listStudyResult == nil {
		listStudyResult = s.client.GetStudyResult(studentId)
		_, err := s.cacheService.TrySetStudyResult(ctx, studentId+"_"+strconv.Itoa(semester), listStudyResult)
		if err != nil {
			result.err = errors.New(common.SET_STUDY_RESULT_FAIL_REDIS)
			c <- result
			return
		}
	}

	for _, course := range listStudyResult {
		if course.Result == 3 {
			course := dto.CourseSuggestion{
				CourseId:   course.CourseId,
				CourseName: course.CourseName,
				NumCredits: s.dbConfig.GetCourseConfig(course.CourseId).NumCredits,
				Type:       NOT_PASS,
			}
			*courseSuggestionsFailReasons = append(*courseSuggestionsFailReasons, course)
		}
	}
	result.result = true
	c <- result
}

func (s *registerCourseCheckServiceImp) getCoursesTeachingPlan(ctx context.Context, studentId string, studentInfo *client.StudentInfo, courseSuggestionsTeachingPlan *[]dto.CourseSuggestion, freeCreditInfo *[]dto.FreeCreditInfo, c chan<- chanResult[bool]) {
	result := chanResult[bool]{}
	result.result = false

	listCourse, freeCreditInfoFromDB, err := s.repository.GetListCourseOfTeachingPlan(studentInfo.Falcuty, studentInfo.Speciality, studentInfo.AcademicProgram, studentInfo.SemesterOrder)
	if err != nil {
		result.err = err
		c <- result
		return
	}
	*freeCreditInfo = freeCreditInfoFromDB

	for _, course := range listCourse {
		courseModel := dto.CourseSuggestion{
			Type: TEACHING_PLAN,
		}
		if cf := s.dbConfig.GetCourseConfig(course); cf != nil {
			courseModel.CourseId = course
			courseModel.CourseName = cf.CourseName
			courseModel.NumCredits = s.dbConfig.GetCourseConfig(course).NumCredits
			*courseSuggestionsTeachingPlan = append(*courseSuggestionsTeachingPlan, courseModel)
		}

	}
	result.result = true
	c <- result
}

func (s *registerCourseCheckServiceImp) getMinMaxCreditAsync(ctx context.Context, studentId string, academicProgram string, semester int, minCredit *int, maxCredit *int, c chan<- chanResult[bool]) {
	result := chanResult[bool]{}
	result.result = false

	minMaxCredit, _ := s.cacheService.GetMinMaxCredit(ctx, studentId+"_"+strconv.Itoa(semester))

	if minMaxCredit == nil {
		min, max, err := s.repository.GetMinMaxCredit(studentId, academicProgram, semester)
		*minCredit = min
		*maxCredit = max

		if err != nil {
			result.err = err
			c <- result
			return
		}
		_, err = s.cacheService.TrySetMinMaxCredit(ctx, studentId+"_"+strconv.Itoa(semester), []int{*minCredit, *maxCredit})

		if err != nil {
			result.err = errors.New(common.SET_MIN_MAX_CREDIT_FAIL_REDIS)
			c <- result
			return
		}
	} else {
		*minCredit = minMaxCredit[0]
		*maxCredit = minMaxCredit[1]
	}
	result.result = true
	c <- result
}
