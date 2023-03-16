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

func (c *client) GetStudentStatus(studentId int) int {
	// Get student status from core service
	//http.Get("")
	if studentId == 1915982 || studentId == 1915983 || studentId == 1914698 {
		return 1
	} else {
		return 0
	}
}

func (c *client) GetStudyResult(studentId int) []CourseResult {

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
