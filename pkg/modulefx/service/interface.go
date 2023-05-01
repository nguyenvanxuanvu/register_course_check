package service

import (
	"context"

	"github.com/nguyenvanxuanvu/register_course_check/pkg/dto"
)

type RegisterCourseCheckService interface {
	Check(ctx context.Context, req *dto.CheckRequestDTO) (*dto.CheckResponseDTO, error)
	Suggestion(ctx context.Context, req *dto.SuggestionRequestDTO) (*dto.SuggestionResponseDTO, error)
	UpdateCourseCondition(ctx context.Context, req []dto.CourseConditionConfig) (bool, error)
}
