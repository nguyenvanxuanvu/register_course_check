package dbconfig

import (
	"register_course_check/pkg/dto"
	

	
)

type DBConfig interface {
	GetSubjectConfig(subjectId string) (*dto.SubjectConfig)
}
