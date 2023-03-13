package client

type StudentInfo struct {
	StudentId int `json:"studentId"`
	StudentName string `json:"studentName"`
	Falcuty string     `json:"falcuty"`
	AcademicProgram  string  `json:"academicProgram"`
	StudyResults []CourseResult `json:"studyResult"`

}

type CourseResult struct {
	CourseId string `json:"courseId"`
	CourseName string `json:"courseName"`
	Result  int 	`json:"result"`              // 1 : Dat   2: Dang hoc   3: Khong dat   4: Chua hoc
}