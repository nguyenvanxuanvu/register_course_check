package dto

import (
	"database/sql/driver"
	"encoding/json"
)

type CourseConfig struct {
	Id                    string `db:"id"`
	CourseName            string `db:"course_name"`
	NumCredits            int    `db:"num_credits"`
	Faculty               string `db:"faculty"`
	CourseConditionConfig *CourseConditionConfig
}

type CourseConditionConfig struct {
	CourseId  string           `db:"course_id"`
	Condition *CourseCondition `db:"condition"` // json type - object struct
}

type CourseCondition struct {
	Op     string               `db:"op,omitempty"`
	Course *CourseConditionInfo `db:"course,omitempty"`
	Leaves []*CourseCondition   `db:"leaves,omitempty"`
}

// Type
// 1: Tien quyet   2: Hoc truoc  3: Song hanh

type CourseConditionInfo struct {
	CourseDesId string `db:"courseDesId,omitempty"`
	Type        int    `db:"type,omitempty"`
}

type FreeCreditInfo struct {
	Group string `db:"group,omitempty"`
	Nums int `db:"nums,omitempty"`
}


// Value implements the driver.Valuer interface
func (f CourseCondition) Value() (driver.Value, error) {
	return json.Marshal(f)
}

// Scan implements the sql.Scanner interface
func (f *CourseCondition) Scan(value interface{}) error {
	var data = []byte(value.([]uint8))
	return json.Unmarshal(data, &f)
}
