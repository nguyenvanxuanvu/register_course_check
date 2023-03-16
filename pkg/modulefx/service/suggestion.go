package service

import (
	"context"
	
	"github.com/nguyenvanxuanvu/register_course_check/pkg/dto"
)

func (s *registerCourseCheckServiceImp) Suggestion(ctx context.Context, req *dto.SuggestionRequestDTO) (*dto.SuggestionResponseDTO, error) {
	return &dto.SuggestionResponseDTO{
		MinCredit: 2,
		MaxCredit: 3,
	}, nil
}