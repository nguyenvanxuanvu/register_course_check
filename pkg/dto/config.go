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
	Condition *CourseCondition `db:"course_condition"` // json type - object struct
}

type CourseCondition struct {
	Op     string               `json:"op,omitempty"`
	Course *CourseConditionInfo `json:"course,omitempty"`
	Leaves []*CourseCondition   `json:"leaves,omitempty"`
}

// Type
// 1: Tien quyet   2: Hoc truoc  3: Song hanh

type CourseConditionInfo struct {
	CourseDesId string `json:"courseDesId,omitempty"`
	Type        int    `json:"type,omitempty"`
}

type FreeCreditInfo struct {
	Group string `db:"group,omitempty"`
	Nums  int    `db:"nums,omitempty"`
}

// Value implements the driver.Valuer interface
func (f CourseCondition) Value() (driver.Value, error) {
	return json.Marshal(f)
}

// Scan implements the sql.Scanner interface
func (f *CourseCondition) Scan(value interface{}) error {

	var data []byte
	switch value.(type) {
	case []uint8:
		data = []byte(value.([]uint8))
	case string:
		data = []byte(value.(string))
	}

	return json.Unmarshal(data, &f)
}
