package repository

import (

	"register_course_check/pkg/dto"
)

type ConfigRepository interface {
	GetSubjectConfigs() ([]*dto.SubjectConfig, error)
	GetSubjectConditionConfigs() (map[string][]*dto.SubjectCondtionConfig, error)
	
}

type Repository interface {
	GetMinMaxCredit(academicProgram string, semester int) (int, int)
}


