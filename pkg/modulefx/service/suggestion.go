package service

import (
	"context"
	"errors"

	"github.com/nguyenvanxuanvu/register_course_check/pkg/common"
	"github.com/nguyenvanxuanvu/register_course_check/pkg/dto"
	"github.com/nguyenvanxuanvu/register_course_check/pkg/modulefx/client"
)

func (s *registerCourseCheckServiceImp) Suggestion(ctx context.Context, req *dto.SuggestionRequestDTO) (*dto.SuggestionResponseDTO, error) {
	studentId := int(req.StudentId)
	var listStudyResult []client.CourseResult
	listStudyResult, _ = s.cacheService.GetStudyResult(ctx, studentId)
	if listStudyResult == nil {
		listStudyResult = s.client.GetStudyResult(int(req.StudentId))
		_ , err := s.cacheService.TrySetStudyResult(ctx, studentId, listStudyResult)
		if err != nil {
			return nil, errors.New(common.SET_STUDY_RESULT_FAIL_REDIS)
		}
	}
	var courseSuggestions []dto.CourseSuggestion
	for _, course := range listStudyResult{
		if course.Result == 3 {
			course := dto.CourseSuggestion{
				CourseId: course.CourseId,
				CourseName: course.CourseName,
				Type: 1,
			}
			courseSuggestions = append(courseSuggestions, course)
		}
	}

	
	
	
	return &dto.SuggestionResponseDTO{
		Courses: courseSuggestions,
		MinCredit: 2,
		MaxCredit: 3,
	}, nil
}