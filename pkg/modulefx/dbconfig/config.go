package dbconfig

import (
	"github.com/nguyenvanxuanvu/register_course_check/pkg/dto"
	"github.com/nguyenvanxuanvu/register_course_check/pkg/modulefx/repository"
)

type dbConfigObj struct {
	courseConfigs      map[string]*dto.CourseConfig
}

func NewDBConfig(repo repository.ConfigRepository) (DBConfig, error) {
	
	courseConfigs, err := repo.GetCourseConfigs()
	if err != nil {
		return nil, err
	}

	courseConditionConfigs, err := repo.GetCourseConditionConfigs()
	if err != nil {
		return nil, err
	}
	


	courseConfigMap := make(map[string]*dto.CourseConfig)

	for _, courseConfig := range courseConfigs {
		
		courseConfig.CourseConditionConfig = &dto.CourseConditionConfig{}
		condition := courseConditionConfigs[courseConfig.Id]
		courseConfig.CourseConditionConfig = condition
		courseConfigMap[courseConfig.Id] = courseConfig
	}

	
	return &dbConfigObj{
		courseConfigMap,
		
	}, nil

}


func (c *dbConfigObj) GetCourseConfig(courseId string) (*dto.CourseConfig) {
	return c.courseConfigs[courseId]
}


