package repository

import (

	"register_course_check/pkg/dto"
)

type ConfigRepository interface {
	GetSubjectConfigs() ([]*dto.SubjectConfig, error)
	GetSubjectConditionConfigs() (map[string][]*dto.SubjectCondtionConfig, error)
	
}

type Repository interface {
	GetStudentStatus(studentId int) (int)
	GetMinCredit(academicProgram string, semester int) (int)
	GetListDoneCourse(studentId int) []string
}


