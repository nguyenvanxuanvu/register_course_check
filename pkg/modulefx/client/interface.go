package client



type Client interface {
	GetStudentInfo(studentId string) *StudentInfo
	GetStudyResult(studentId string) []CourseResult
}

