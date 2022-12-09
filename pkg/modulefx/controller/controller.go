package controller

import (
	"register_course_check/pkg/modulefx/service"
)

type controllerImpl struct {
	registerCourseCheckService service.RegisterCourseCheckService
}

func NewController(registerCourseCheckService service.RegisterCourseCheckService) Controller {
	return &controllerImpl{
		registerCourseCheckService: registerCourseCheckService,
	}
}
