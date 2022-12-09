package main

import (
	
	"log"
	"os"
	"go.uber.org/fx"
	"register_course_check/httpserver"
	"register_course_check/mysql"
	"register_course_check/config"
	
	
	
	
	"register_course_check/pkg/modulefx"
	
)

var GitCommit string

func printInfo() {
	log.Println("Environment:", os.Getenv("appenv"))
	log.Println("Git Commit:", GitCommit)
}

func main() {
	printInfo()

	app := fx.New(
		
		config.Module,
		httpserver.Module,
		mysql.Module,
		modulefx.Module,
		
		
		
	)
	app.Run()
}
