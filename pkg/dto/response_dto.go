package dto


type CheckResponseDTO struct {
	Status  string `json:"status"`
	StudentStatus string `json:"studentStatus"`
	CourseChecks []*CourseCheck `json:"courseChecks"`
	CheckMinCreditResult MinMaxCredit `json:"checkMinCreditResult"`
	CheckMaxCreditResult MinMaxCredit `json:"checkMaxCreditResult"`
}

type MinMaxCredit struct {
	CheckResult string `json:"checkResult"`
	CurrentRegister int `json:"currentRegister,omitempty"`
	Config int `json:"config,omitempty"`
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


type SuggestionResponseDTO struct {
	Courses []CourseSuggestion `json:"courses"`
	MinCredit int `json:"minCredit"`
	MaxCredit int `json:"maxCredit"`
}

type CourseSuggestion struct {
	CourseId string  `json:"courseId"`
	CourseName string  `json:"courseName"` 
	Type int `json:"type"`        // 1: Chua dat   2: Chuong trinh hoc
}