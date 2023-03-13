package service

import (
	"register_course_check/pkg/modulefx/client"
	db_config "register_course_check/pkg/modulefx/dbconfig"
	"register_course_check/pkg/modulefx/repository"
)

type registerCourseCheckServiceImp struct {
	dbConfig             db_config.DBConfig
	repository			 repository.Repository
	client				 client.Client
}

func NewRegisterCourseCheckService(
	dbConfig db_config.DBConfig,
	repository repository.Repository,
	client  client.Client,
	
) RegisterCourseCheckService {
	return &registerCourseCheckServiceImp{
		dbConfig,
		repository,
		client,
	}
}


