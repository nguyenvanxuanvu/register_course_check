package client



type Client interface {
	GetStudentStatus(studentId int) int
	GetListDoneCourse(studentId int) []string
}

