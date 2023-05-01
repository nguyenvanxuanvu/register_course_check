package controller

import (
	"context"

	"github.com/nguyenvanxuanvu/register_course_check/pkg/dto"
)

type Controller interface {
	Check(ctx context.Context, request *dto.CheckRequestDTO) (*dto.CheckResponseDTO, error)
	Suggestion(ctx context.Context, request *dto.SuggestionRequestDTO) (*dto.SuggestionResponseDTO, error)
	UpdateCourseCondition(ctx context.Context, request []dto.CourseConditionConfig) (bool, error)
}
