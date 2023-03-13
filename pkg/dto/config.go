package dto

import (
	"database/sql/driver"
	"encoding/json"
)

type SubjectConfig struct {
	Id                     string `db:"id"`
	SubjectName            string `db:"subject_name"`
	NumCredits             int    `db:"num_credits"`
	Faculty                string `db:"faculty"`
	SubjectConditionConfig *SubjectConditionConfig
}

type SubjectConditionConfig struct {
	SubjectId string           `db:"subject_id"`
	Condition SubjectCondition `db:"condition"` // json type - object struct
}




// subject condition type json
// subjectId - Data has the format like "CO2-number"   number: 1: Tien quyet   2: Hoc truoc  3: Song hanh

type SubjectCondition struct {
	Left  *SubjectCondition  `db:"left,omitempty"`
	Right *SubjectCondition  `db:"right,omitempty"`
	Data  string 			 `db:"data,omitempty"`
}









// Value implements the driver.Valuer interface
func (f SubjectCondition) Value() (driver.Value, error) {
	return json.Marshal(f)
}

// Scan implements the sql.Scanner interface
func (f *SubjectCondition) Scan(value interface{}) error {
	var data = []byte(value.([]uint8))
	return json.Unmarshal(data, &f)
}
