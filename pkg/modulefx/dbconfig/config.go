package dbconfig

import (
	
	"register_course_check/pkg/dto"
	"register_course_check/pkg/modulefx/repository"
)

type dbConfigObj struct {
	subjectConfigs      map[string]*dto.SubjectConfig
}

func NewDBConfig(repo repository.ConfigRepository) (DBConfig, error) {
	
	subjectConfigs, err := repo.GetSubjectConfigs()
	if err != nil {
		return nil, err
	}
	

	subjectConditionConfigs, err := repo.GetSubjectConditionConfigs()
	if err != nil {
		return nil, err
	}

	subjectConfigMap := make(map[string]*dto.SubjectConfig)

	for _, subjectConfig := range subjectConfigs {
		
		subjectConfig.SubjectConditionConfig = []*dto.SubjectCondtionConfig{}
		conditions := subjectConditionConfigs[subjectConfig.Id]
			for  _,condition  := range conditions{
				
				subjectConfig.SubjectConditionConfig = append(subjectConfig.SubjectConditionConfig, condition)
			}
			
			subjectConfigMap[subjectConfig.Id] = subjectConfig
		

	}

	
	return &dbConfigObj{
		subjectConfigMap,
		
	}, nil

}


func (c *dbConfigObj) GetSubjectConfig(subjectId string) *dto.SubjectConfig {
	return c.subjectConfigs[subjectId]
}

