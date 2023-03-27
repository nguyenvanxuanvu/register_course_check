package repository

import (
	"github.com/nguyenvanxuanvu/register_course_check/pkg/dto"

	"github.com/jmoiron/sqlx"
)

type configRepository struct {
	db *sqlx.DB
}

func NewConfigRepository(db *sqlx.DB) ConfigRepository {
	return &configRepository{db: db}
}

const COURSE_TABLE = "course"
const SELECT_COURSE_CONFIG = "SELECT `id`,`course_name`,`num_credits`,`faculty` FROM `" + COURSE_TABLE + "`"

func (r *configRepository) GetCourseConfigs() ([]*dto.CourseConfig, error) {
	rows, err := r.db.Queryx(SELECT_COURSE_CONFIG)
	if err != nil {
		return nil, err
	}
	courseConfigs := []*dto.CourseConfig{}

	for rows.Next() {
		courseConfig := &dto.CourseConfig{}
		err = rows.StructScan(courseConfig)
		if err != nil {
			return nil, err
		}
		courseConfigs = append(courseConfigs, courseConfig)
	}
	return courseConfigs, nil
}

const COURSE_CONDITION_TABLE = "course_condition"
const SELECT_COURSE_CONDITION_CONFIG = "SELECT `course_id`,`condition` FROM `" + COURSE_CONDITION_TABLE + "`"

func (r *configRepository) GetCourseConditionConfigs() (map[string]*dto.CourseConditionConfig, error) {
	rows, err := r.db.Queryx(SELECT_COURSE_CONDITION_CONFIG)
	if err != nil {
		return nil, err
	}
	courseConditionConfigs := map[string]*dto.CourseConditionConfig{}
	for rows.Next() {
		conditionConfig := &dto.CourseConditionConfig{}

		err = rows.StructScan(&conditionConfig)

		if err != nil {
			return nil, err
		}

		courseConditionConfigs[conditionConfig.CourseId] = conditionConfig
	}

	return courseConditionConfigs, nil
}
