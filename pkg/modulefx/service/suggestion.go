package service

import (
	"context"
	"errors"

	"github.com/nguyenvanxuanvu/register_course_check/pkg/common"
	"github.com/nguyenvanxuanvu/register_course_check/pkg/dto"
	"github.com/nguyenvanxuanvu/register_course_check/pkg/modulefx/client"
)

const (
	NOT_PASS = 1
	TEACHING_PLAN = 2
)
func (s *registerCourseCheckServiceImp) Suggestion(ctx context.Context, req *dto.SuggestionRequestDTO) (*dto.SuggestionResponseDTO, error) {
	studentId := int(req.StudentId)
	var listStudyResult []client.CourseResult
	listStudyResult, _ = s.cacheService.GetStudyResult(ctx, studentId)
	if listStudyResult == nil {
		listStudyResult = s.client.GetStudyResult(int(req.StudentId))
		_, err := s.cacheService.TrySetStudyResult(ctx, studentId, listStudyResult)
		if err != nil {
			return nil, errors.New(common.SET_STUDY_RESULT_FAIL_REDIS)
		}
	}
	var courseSuggestions []dto.CourseSuggestion
	for _, course := range listStudyResult {
		if course.Result == 3{
			course := dto.CourseSuggestion{
				CourseId:   course.CourseId,
				CourseName: course.CourseName,
				NumCredits: s.dbConfig.GetCourseConfig(course.CourseId).NumCredits,
				Type:       NOT_PASS,
			}
			courseSuggestions = append(courseSuggestions, course)
		}
	}

	var studentInfo *client.StudentInfo
	studentInfo, _ = s.cacheService.GetStudentInfo(ctx, studentId)
	if studentInfo == nil {
		studentInfo = s.client.GetStudentInfo(int(req.StudentId))
		_, err := s.cacheService.TrySetStudentInfo(ctx, studentId, studentInfo)
		if err != nil {
			return nil, errors.New(common.SET_STUDENT_INFO_FAIL_REDIS)
		}
	}

	listCourse, freeCreditInfo, err := s.repository.GetListCourseOfTeachingPlan(studentInfo.Falcuty, studentInfo.Speciality, studentInfo.AcademicProgram, studentInfo.SemesterOrder)
	if err != nil {
		return nil, err
	}

	for _, course := range listCourse {
		courseModel := dto.CourseSuggestion{
			Type: TEACHING_PLAN,
		}
		if cf := s.dbConfig.GetCourseConfig(course); cf != nil {
			courseModel.CourseId = course
			courseModel.CourseName = cf.CourseName
			courseModel.NumCredits = s.dbConfig.GetCourseConfig(course).NumCredits
			courseSuggestions = append(courseSuggestions, courseModel)
		}

		

	}
	var minCredit, maxCredit int
	minMaxCredit,_ := s.cacheService.GetMinMaxCredit(ctx, studentId)
	if minMaxCredit == nil {
		minCredit, maxCredit, err = s.repository.GetMinMaxCredit(studentId, studentInfo.AcademicProgram, int(req.Semester))
		if err != nil{
			return nil, err
		}
		_, err := s.cacheService.TrySetMinMaxCredit(ctx, studentId, []int{minCredit, maxCredit})
		if err != nil {
			return nil, errors.New(common.SET_MIN_MAX_CREDIT_FAIL_REDIS)
		}
	} else {
		minCredit = minMaxCredit[0]
		maxCredit = minMaxCredit[1]
	}
	
	return &dto.SuggestionResponseDTO{
		Courses:          courseSuggestions,
		HintOfFreeCredit: freeCreditInfo,
		MinCredit:        minCredit,
		MaxCredit:        maxCredit,
	}, nil
}
