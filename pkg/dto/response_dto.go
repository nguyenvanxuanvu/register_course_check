package dto


type CheckResponseDTO struct {
	Status  string `json:"status"`
	StudentStatus string `json:"studentStatus"`
	SubjectChecks []*SubjectCheck `json:"subjectChecks"`
	CheckMinCreditResult string `json:"checkMinCreditResult"`
}

type SubjectCheck struct {
	SubjectId string  `json:"subjectId"`
	SubjectName string  `json:"subjectName"`
	CheckResult string `json:"checkResult"`
	FailReasons []*Reason `json:"failReasons"`
}

type Reason struct {
	SubjectDesId string `json:"subjectDesId"`
	ConditionType int `json:"conditionType"`         //  1: TQ   2: HT   3:SH
}
