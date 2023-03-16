package main

import (
	"log"
	"os"
	"github.com/nguyenvanxuanvu/register_course_check/config"
	"github.com/nguyenvanxuanvu/register_course_check/httpserver"
	"github.com/nguyenvanxuanvu/register_course_check/mysql"
	"github.com/nguyenvanxuanvu/register_course_check/redis"

	"go.uber.org/fx"

	"github.com/nguyenvanxuanvu/register_course_check/pkg/modulefx"
)

var GitCommit string

func printInfo() {
	log.Println("Environment:", os.Getenv("appenv"))
}

func main() {
	printInfo()

	app := fx.New(

		config.Module,
		httpserver.Module,
		mysql.Module,
		modulefx.Module,
		redis.Module,
	)
	app.Run()
}
