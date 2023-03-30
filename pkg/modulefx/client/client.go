package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type client struct {
}

func NewClient() Client {
	return &client{}
}

func (c *client) GetStudentInfo(studentId string) *StudentInfo {
	// Get student status from core service
	//http.Get("")
	jsonFile, err := os.Open("pkg/modulefx/client/student.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var students Students

	json.Unmarshal(byteValue, &students)

	// Get student study result from core service

	for _, student := range students.Students {
		if student.StudentId == studentId {
			return &StudentInfo{
				StudentStatus: student.StudentStatus,
				AcademicProgram: student.AcademicProgram,
				Falcuty: student.Falcuty,
				Speciality: student.Speciality,
				SemesterOrder: student.SemesterOrder,
			}
		}
	}
	return nil

}

func (c *client) GetStudyResult(studentId string) []CourseResult {

	jsonFile, err := os.Open("pkg/modulefx/client/student.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var students Students

	json.Unmarshal(byteValue, &students)

	// Get student study result from core service

	for _, student := range students.Students {
		if student.StudentId == studentId {
			return student.StudyResults
		}
	}
	return nil

}
