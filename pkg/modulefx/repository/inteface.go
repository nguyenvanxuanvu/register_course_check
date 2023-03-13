package repository

import (

	"register_course_check/pkg/dto"
)

type ConfigRepository interface {
	GetSubjectConfigs() ([]*dto.SubjectConfig, error)
	GetSubjectConditionConfigs() (map[string]*dto.SubjectConditionConfig, error)
	
}

type Repository interface {
	GetMinMaxCredit(academicProgram string, semester int) (int, int)
}


