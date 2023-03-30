package client


type Students struct {
    Students []Student `json:"students"`
}
type Student struct {
	StudentId string `json:"studentId"`
	StudentName string `json:"studentName"`
	StudentStatus int `json:"studentStatus"`
	Falcuty string     `json:"falcuty"`
	AcademicProgram  string  `json:"academicProgram"`
	Speciality  string      `json:"speciality"`
	SemesterOrder int `json:"semesterOrder"`
	StudyResults []CourseResult `json:"studyResults"`
}

type StudentInfo struct {
	StudentStatus int `json:"studentStatus"`
	Falcuty string     `json:"falcuty"`
	AcademicProgram  string  `json:"academicProgram"`
	Speciality  string      `json:"speciality"`
	SemesterOrder int `json:"semesterOrder"`
} 

type CourseResult struct {
	CourseId string `json:"courseId"`
	CourseName string `json:"courseName"`
	Result  int 	`json:"result"`              // 1 : Dat   2: Dang hoc   3: Khong dat
}