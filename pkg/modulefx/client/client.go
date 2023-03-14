package client

type client struct {
}

func NewClient() Client {
	return &client{}
}

func (c *client) GetStudentStatus(studentId int) int {
	// Get student status from core service
	//http.Get("")
	if studentId == 1915982 {
		return 1
	} else {
		return 0
	}
}

func (c *client) GetStudyResult(studentId int) []CourseResult {
	// Get student study result from core service
	if studentId == 1915982 {
		studentInfo := &StudentInfo{
			StudentId:       1915982,
			StudentName:     "Nguyễn Văn Xuân Vũ",
			Falcuty:         "KHMT",
			AcademicProgram: "DT",
			StudyResults: []CourseResult{
				{
					CourseId:   "CO1",
					CourseName: "Công nghệ phần mềm",
					Result:     1,
				},
				{
					CourseId:   "CO4",
					CourseName: "Lập trình web",
					Result:     3,
				},
				{
					CourseId:   "CO3",
					CourseName: "Lập trình web",
					Result:     1,
				},
			},
		}

		return studentInfo.StudyResults

	} else {
		return nil
	}
}
