package repository

import (
	"log"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

const STUDENT_TABLE = "student"
const SELECT_STUDENT_CONFIG = "SELECT student_status FROM " + STUDENT_TABLE + " where id = ?;"

func (r *repository) GetStudentStatus(studentId int) (int) {
	
	var status int 
	rows, err := r.db.Queryx(SELECT_STUDENT_CONFIG, studentId)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		rows.Scan(&status)
	}
	return status
}

const MIN_CREDIT_TABLE = "min_credit"
const SELECT_MIN_CREDIT_CONFIG = "SELECT min_credit FROM " + MIN_CREDIT_TABLE + " WHERE academic_program = ? AND semester = ?;"

func (r *repository) GetMinCredit(academinProgram string, semester int) (int) {
	var numCredits int 
	rows, err := r.db.Queryx(SELECT_MIN_CREDIT_CONFIG, academinProgram, semester)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		rows.Scan(&numCredits)
	}
	
	return numCredits
}

const RESULT_COURSE_TABLE = "result"
const SELECT_DONE_COURSE_CONFIG = "SELECT subject FROM " + RESULT_COURSE_TABLE + " WHERE student_id = ? AND result = 1;"

func (r *repository) GetListDoneCourse(studentId int) []string {
	rows, err := r.db.Queryx(SELECT_DONE_COURSE_CONFIG, studentId)
	
	if err != nil {
		return nil
	}
	var doneCourseList []string 

	for rows.Next() {
		var doneCourse string
		err = rows.Scan(&doneCourse)
		if err != nil {
			return nil
		}
		doneCourseList = append(doneCourseList, doneCourse)
	}
	return doneCourseList
}





