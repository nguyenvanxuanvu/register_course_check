package repository

import (

	"github.com/nguyenvanxuanvu/register_course_check/pkg/dto"
)

type ConfigRepository interface {
	GetCourseConfigs() ([]*dto.CourseConfig, error)
	GetCourseConditionConfigs() (map[string]*dto.CourseConditionConfig, error)
	
}

type Repository interface {
	GetMinMaxCredit(academicProgram string, semester int) (int, int)
}


