package dto



type CheckRequestDTO struct {
	StudentId uint64 `json:"studentId"`
	AcademicProgram string `json:"academicProgram"`
	Semester uint64	`json:"semester"`
	RegisterSubjects []*Subject `json:"registerSubjects"`
}

type Subject struct {
	SubjectId string `json:"subjectId"`
}

