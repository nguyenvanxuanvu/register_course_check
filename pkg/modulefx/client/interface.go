package client



type Client interface {
	GetStudentStatus(studentId int) int
	GetStudyResult(studentId int) []CourseResult
}

