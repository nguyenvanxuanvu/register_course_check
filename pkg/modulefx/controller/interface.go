package controller

import (
	"context"

	"register_course_check/pkg/dto"
)

type Controller interface {
	Check(ctx context.Context, request *dto.CheckRequestDTO) (*dto.CheckResponseDTO, error)
}
