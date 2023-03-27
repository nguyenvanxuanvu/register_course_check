package common

const (
	NOT_FOUND_STUDENT_STATUS                    = "NOT_FOUND_STUDENT_STATUS"
	NOT_FOUND_MIN_CREDIT_CONFIG					= "NOT_FOUND_MIN_CREDIT_CONFIG"
	NOT_FOUND_MAX_CREDIT_CONFIG					= "NOT_FOUND_MAX_CREDIT_CONFIG"
	MIN_MAX_CONFIG_WRONG						= "MIN_MAX_CONFIG_WRONG"
	NOT_FOUND_COURSE_ID							= "NOT_FOUND_COURSE_ID"
	NOT_FOUND_COURSE_REGISTER					= "NOT_FOUND_COURSE_REGISTER"
	DUPLICATE_COURSE_REGISTER					= "DUPLICATE_COURSE_REGISTER"
	SET_STUDENT_INFO_FAIL_REDIS					= "SET_STUDENT_INFO_FAIL_REDIS"
	SET_STUDY_RESULT_FAIL_REDIS				    = "SET_STUDY_RESULT_FAIL_REDIS"
)


var ErrToDescription = map[string]string{
    NOT_FOUND_STUDENT_STATUS : "Không tìm thấy trạng thái sinh viên",
	NOT_FOUND_MIN_CREDIT_CONFIG: "Không tìm thấy thông tín số tín chỉ tối thiểu",
	NOT_FOUND_MAX_CREDIT_CONFIG: "Không tìm thấy thông tín số tín chỉ tối đa",
	MIN_MAX_CONFIG_WRONG: "Sai cấu hình số tín chỉ tối thiểu, tối đa",
	NOT_FOUND_COURSE_ID: "Không tìm thấy môn học",
	NOT_FOUND_COURSE_REGISTER: "Cần điền thông tin môn học đăng ký",
	DUPLICATE_COURSE_REGISTER: "Trùng môn học đăng ký",
	SET_STUDENT_INFO_FAIL_REDIS: "Lỗi hệ thống",
	SET_STUDY_RESULT_FAIL_REDIS: "Lỗi hệ thống",

}