package dbconfig

import (
	"register_course_check/pkg/dto"
	

	
)

type DBConfig interface {
	GetCourseConfig(courseId string) (*dto.CourseConfig)
}
