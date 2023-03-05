package modulefx

import (
	"register_course_check/pkg/modulefx/client"
	"register_course_check/pkg/modulefx/controller"
	"register_course_check/pkg/modulefx/dbconfig"
	"register_course_check/pkg/modulefx/repository"
	"register_course_check/pkg/modulefx/router"
	"register_course_check/pkg/modulefx/service"

	"go.uber.org/fx"
)

var Module = fx.Options(
	controller.Module,
	router.Module,
	repository.Module,
	dbconfig.Module,
	service.Module,
	client.Module,
)
