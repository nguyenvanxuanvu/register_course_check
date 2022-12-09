package controller

import (
	"context"
	

	"register_course_check/pkg/dto"
)


func (s *controllerImpl) Check(ctx context.Context, req *dto.CheckRequestDTO) (*dto.CheckResponseDTO, error) {
	return s.registerCourseCheckService.Check(ctx, req)
}
