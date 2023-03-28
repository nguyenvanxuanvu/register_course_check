package dto

type CheckRequestDTO struct {
	StudentId       uint64    `json:"studentId"`
	Semester        uint64    `json:"semester"`
	RegisterCourses []*Course `json:"registerCourses"`
}

type Course struct {
	CourseId string `json:"courseId"`
}

type SuggestionRequestDTO struct {
	StudentId uint64 `json:"studentId"`
	Semester  uint64 `json:"semester"`
}
