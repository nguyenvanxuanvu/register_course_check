package modulefx

import (
	"github.com/nguyenvanxuanvu/register_course_check/pkg/modulefx/authen"
	"github.com/nguyenvanxuanvu/register_course_check/pkg/modulefx/cache"
	"github.com/nguyenvanxuanvu/register_course_check/pkg/modulefx/client"
	"github.com/nguyenvanxuanvu/register_course_check/pkg/modulefx/controller"
	"github.com/nguyenvanxuanvu/register_course_check/pkg/modulefx/dbconfig"
	"github.com/nguyenvanxuanvu/register_course_check/pkg/modulefx/repository"
	"github.com/nguyenvanxuanvu/register_course_check/pkg/modulefx/router"
	"github.com/nguyenvanxuanvu/register_course_check/pkg/modulefx/service"

	"go.uber.org/fx"
)

var Module = fx.Options(
	controller.Module,
	router.Module,
	repository.Module,
	dbconfig.Module,
	service.Module,
	client.Module,
	cache.Module,
	authen.Module,
)
