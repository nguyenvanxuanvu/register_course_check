package repository

import (
	"register_course_check/pkg/dto"

	"github.com/jmoiron/sqlx"
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
const SELECT_SUBJECT_CONDITION_CONFIG = "SELECT `subject_id`,`condition` FROM `" + SUBJECT_CONDITION_TABLE + "`"

func (r *configRepository) GetSubjectConditionConfigs() (map[string]*dto.SubjectConditionConfig, error) {
	rows, err := r.db.Queryx(SELECT_SUBJECT_CONDITION_CONFIG)
	if err != nil {
		return nil, err
	}
	subjectConditionConfigs := map[string]*dto.SubjectConditionConfig{}
	for rows.Next() {
		conditionConfig := &dto.SubjectConditionConfig{}
		err = rows.StructScan(conditionConfig)
		if err != nil {
			return nil, err
		}
		subjectConditionConfigs[conditionConfig.SubjectId] = conditionConfig
	}
	
	return subjectConditionConfigs, nil
}







