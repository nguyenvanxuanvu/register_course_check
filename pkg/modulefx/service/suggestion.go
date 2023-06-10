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
	minMaxCreditChan := make(chan chanResult[bool], 1)

	// get fail course list of student
	var listStudyResult []client.CourseResult
	go s.getFailCoursesAsync(ctx, studentId, semester, &courseSuggestionsFailReasons, &listStudyResult, failCoursesChan)

	// get min max credit config for student
	var minCredit, maxCredit int
	go s.getMinMaxCreditAsync(ctx, studentId, studentInfo.AcademicProgram, semester, &minCredit, &maxCredit, minMaxCreditChan)

	failCoursesRes, minMaxCreditRes := <-failCoursesChan, <-minMaxCreditChan
	if failCoursesRes.err != nil || minMaxCreditRes.err != nil {
		return nil, oneOf(failCoursesRes.err, minMaxCreditRes.err)
	}

	// get list course of teaching plan of student

	var freeCreditInfo []dto.FreeCreditInfo
	var courseSuggestionsTeachingPlan []dto.CourseSuggestion
	s.getCoursesTeachingPlan(ctx, studentId, studentInfo, &courseSuggestionsTeachingPlan, &freeCreditInfo, listStudyResult)

	var lastRTcourseSuggestionTeachingPlan []dto.CourseSuggestion
	for _, courseInTeachingPlan := range courseSuggestionsTeachingPlan {
		flag := false
		for _, course := range courseSuggestionsFailReasons {
			if course.CourseId == courseInTeachingPlan.CourseId {
				flag = true
				break
			}
		}
		if flag {
			continue
		}
		lastRTcourseSuggestionTeachingPlan = append(lastRTcourseSuggestionTeachingPlan, courseInTeachingPlan)

	}
	// FOR UNIQUE
	var lastForUniquecourseSuggestionTeachingPlan []dto.CourseSuggestion
	keys := make(map[string]bool)

	for _, suggestion := range lastRTcourseSuggestionTeachingPlan{
		if _, value := keys[suggestion.CourseId]; !value {
            keys[suggestion.CourseId] = true
            lastForUniquecourseSuggestionTeachingPlan = append(lastForUniquecourseSuggestionTeachingPlan, suggestion)
        }
	}

	failReasons = append(courseSuggestionsFailReasons, lastForUniquecourseSuggestionTeachingPlan...)
	return &dto.SuggestionResponseDTO{
		Courses:          failReasons,
		HintOfFreeCredit: freeCreditInfo,
		MinCredit:        minCredit,
		MaxCredit:        maxCredit,
	}, nil
}

func (s *registerCourseCheckServiceImp) getFailCoursesAsync(ctx context.Context, studentId string, semester int, courseSuggestionsFailReasons *[]dto.CourseSuggestion, listStudyResultReturn *[]client.CourseResult, c chan<- chanResult[bool]) {
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
	// list study return
	*listStudyResultReturn = listStudyResult

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

func (s *registerCourseCheckServiceImp) getCoursesTeachingPlan(ctx context.Context, studentId string, studentInfo *client.StudentInfo, courseSuggestionsTeachingPlan *[]dto.CourseSuggestion, freeCreditInfo *[]dto.FreeCreditInfo, listStudyResult []client.CourseResult) error {

	listCourse, freeCreditInfoFromDB, err := s.repository.GetListCourseOfTeachingPlan(studentInfo.Falcuty, studentInfo.Speciality, studentInfo.AcademicProgram, studentInfo.SemesterOrder)
	if err != nil {
		return err
	}
	*freeCreditInfo = freeCreditInfoFromDB

	// loai mon da dat
	var listCourseForCheck []string
	var listCourseNotForCheck []string
	for _, course := range listCourse {
		flag := false
		for _, result := range listStudyResult {
			if result.CourseId == course && result.Result == 1 {
				flag = true
				break

			}

		}
		if flag {
			continue
		}
		if s.dbConfig.GetCourseConfig(course).CourseConditionConfig != nil {
			listCourseForCheck = append(listCourseForCheck, course)
		} else {
			listCourseNotForCheck = append(listCourseNotForCheck, course)
		}

	}
	for _, course := range listCourseNotForCheck {
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

	var courseCheckResults []*dto.CourseCheck

	s.checkListCourseConditionForSuggestion(ctx, &courseCheckResults, listCourseForCheck, studentId, listStudyResult)

	for _, course := range courseCheckResults {
		courseModel := dto.CourseSuggestion{
			Type: TEACHING_PLAN,
		}
		if cf := s.dbConfig.GetCourseConfig(course.CourseId); cf != nil {
			if course.CheckResult == PASS {
				courseModel.CourseId = course.CourseId
				courseModel.CourseName = cf.CourseName
				courseModel.NumCredits = s.dbConfig.GetCourseConfig(course.CourseId).NumCredits
				*courseSuggestionsTeachingPlan = append(*courseSuggestionsTeachingPlan, courseModel)
			} else {
				for _, fail := range course.FailReasons {
					courseModel = dto.CourseSuggestion{
						Type: TEACHING_PLAN,
					}
					courseModel.CourseId = fail.CourseDesId
					courseModel.CourseName = fail.CourseDesName
					courseModel.NumCredits = s.dbConfig.GetCourseConfig(fail.CourseDesId).NumCredits
					*courseSuggestionsTeachingPlan = append(*courseSuggestionsTeachingPlan, courseModel)

				}
			}

		}

	}
	return nil
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
