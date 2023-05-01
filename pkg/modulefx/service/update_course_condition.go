package service

import (
	"context"
	"errors"

	"github.com/nguyenvanxuanvu/register_course_check/pkg/common"
	"github.com/nguyenvanxuanvu/register_course_check/pkg/dto"
)

func (s *registerCourseCheckServiceImp) UpdateCourseCondition(ctx context.Context, req []dto.CourseConditionConfig) (bool, error) {
	_, err := s.repository.UpdateCourseCondition(req)
	if err != nil {
		return false, errors.New(common.UPDATE_COURSE_CONDITION_FAIL)
	}
	return true, nil
}
