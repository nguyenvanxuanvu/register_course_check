package controller

import (
	"context"
	

	"github.com/nguyenvanxuanvu/register_course_check/pkg/dto"
)


func (s *controllerImpl) Check(ctx context.Context, req *dto.CheckRequestDTO) (*dto.CheckResponseDTO, error) {
	return s.registerCourseCheckService.Check(ctx, req)
}

func (s *controllerImpl) Suggestion(ctx context.Context, req *dto.SuggestionRequestDTO) (*dto.SuggestionResponseDTO, error) {
	return s.registerCourseCheckService.Suggestion(ctx, req)
}
