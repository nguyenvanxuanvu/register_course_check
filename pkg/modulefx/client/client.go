package client

type client struct {
}

func NewClient() Client {
	return &client{}
}


func (c *client) GetStudentStatus(studentId int) int {
	// Get student status from core service
	//http.Get("")
	return 1
}


func (c *client) GetListDoneCourse(studentId int) []string {
	// Get student info from core service and get list course student had done
	if studentId == 1915982 {
		return []string{
			"CO2",
		}
	}else {
		return []string{}
	}
}

