package dto


type CheckResponseDTO struct {
	Status  string `json:"status"`
	StudentStatus string `json:"studentStatus"`
	CourseChecks []*CourseCheck `json:"courseChecks"`
	CheckMinCreditResult string `json:"checkMinCreditResult"`
	CheckMaxCreditResult string `json:"checkMaxCreditResult"`
}

type CourseCheck struct {
	CourseId string  `json:"courseId"`
	CourseName string  `json:"courseName"`
	CheckResult string `json:"checkResult"`
	FailReasons []*Reason `json:"failReasons"`
}

type Reason struct {
	CourseDesId string `json:"courseDesId"`
	CourseDesName string `json:"courseDesName"`
	ConditionType int `json:"conditionType"`         //  1: TQ   2: HT   3:SH
}
