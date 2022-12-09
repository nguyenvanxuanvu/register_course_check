package service

import (
	"context"

	"register_course_check/pkg/dto"
)

type RegisterCourseCheckService interface {
	Check(ctx context.Context, req *dto.CheckRequestDTO) (*dto.CheckResponseDTO, error)
}
