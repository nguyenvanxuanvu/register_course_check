package client



type Client interface {
	GetStudentInfo(studentId int) *StudentInfo
	GetStudyResult(studentId int) []CourseResult
}

