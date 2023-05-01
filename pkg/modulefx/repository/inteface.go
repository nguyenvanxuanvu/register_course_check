package repository

import (

	"github.com/nguyenvanxuanvu/register_course_check/pkg/dto"
)

type ConfigRepository interface {
	GetCourseConfigs() ([]*dto.CourseConfig, error)
	GetCourseConditionConfigs() (map[string]*dto.CourseConditionConfig, error)
	
}

type Repository interface {
	GetMinMaxCredit(studentId string, academicProgram string, semester int) (int, int, error)
	GetListCourseOfTeachingPlan(faculty string, speciality string, academicProgram string, semester int) ([]string, []dto.FreeCreditInfo, error)
	UpdateCourseCondition(listCourseCondition []dto.CourseConditionConfig) (bool, error)
}


