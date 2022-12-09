package repository

import (
	"github.com/jmoiron/sqlx"
	"register_course_check/pkg/dto"
)

type configRepository struct {
	db *sqlx.DB
}

func NewConfigRepository(db *sqlx.DB) ConfigRepository {
	return &configRepository{db: db}
}

const SUBJECT_TABLE = "subject"
const SELECT_SUBJECT_CONFIG = "SELECT `id`,`subject_name`,`num_credits`,`faculty` FROM `" + SUBJECT_TABLE + "`"

func (r *configRepository) GetSubjectConfigs() ([]*dto.SubjectConfig, error) {
	rows, err := r.db.Queryx(SELECT_SUBJECT_CONFIG)
	if err != nil {
		return nil, err
	}
	subjectConfigs := []*dto.SubjectConfig{}
	
	for rows.Next() {
		subjectConfig := &dto.SubjectConfig{}
		err = rows.StructScan(subjectConfig)
		if err != nil {
			return nil, err
		}
		subjectConfigs = append(subjectConfigs, subjectConfig)
	}
	return subjectConfigs, nil
}



const SUBJECT_CONDITION_TABLE = "subject_condition"
const SELECT_SUBJECT_CONDITION_CONFIG = "SELECT `subject_id`,`subject_des_id`,`condition_type` FROM `" + SUBJECT_CONDITION_TABLE + "`"

func (r *configRepository) GetSubjectConditionConfigs() (map[string][]*dto.SubjectCondtionConfig, error) {
	rows, err := r.db.Queryx(SELECT_SUBJECT_CONDITION_CONFIG)
	if err != nil {
		return nil, err
	}
	subjectConditionConfigs := map[string][]*dto.SubjectCondtionConfig{}
	for rows.Next() {
		conditionConfig := &dto.SubjectCondtionConfig{}
		err = rows.StructScan(conditionConfig)
		if err != nil {
			return nil, err
		}
		subjectConditionConfigs[conditionConfig.SubjectId] = append(subjectConditionConfigs[conditionConfig.SubjectId], conditionConfig)
	}
	return subjectConditionConfigs, nil
}




