package dto

import (
	"database/sql/driver"
	"encoding/json"
)

type CourseConfig struct {
	Id                     string `db:"id"`
	CourseName            string `db:"course_name"`
	NumCredits             int    `db:"num_credits"`
	Faculty                string `db:"faculty"`
	CourseConditionConfig *CourseConditionConfig
}

type CourseConditionConfig struct {
	CourseId string           `db:"course_id"`
	Condition CourseCondition `db:"condition"` // json type - object struct
}




// course condition type json
// courseId - Data has the format like "CO2-number"   number: 1: Tien quyet   2: Hoc truoc  3: Song hanh

type CourseCondition struct {
	Left  *CourseCondition  `db:"left,omitempty"`
	Right *CourseCondition  `db:"right,omitempty"`
	Data  string 			 `db:"data,omitempty"`
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
