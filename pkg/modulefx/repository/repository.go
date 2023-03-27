package repository

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/nguyenvanxuanvu/register_course_check/pkg/dto"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}



const MIN_MAX_CREDIT_TABLE = "min_max_credit"
const SELECT_MIN_CREDIT_CONFIG = "SELECT min_credit, max_credit FROM " + MIN_MAX_CREDIT_TABLE + " WHERE academic_program = ? AND semester = ?;"

func (r *repository) GetMinMaxCredit(academinProgram string, semester int) (int,int) {
	var numMinCredits, numMaxCredits int 
	rows, err := r.db.Queryx(SELECT_MIN_CREDIT_CONFIG, academinProgram, semester)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		rows.Scan(&numMinCredits, &numMaxCredits)
	}
	
	return numMinCredits, numMaxCredits
}

const TEACHING_PLAN_TABLE = "teaching_plan"
const SELECT_TEACHING_PLAN_CONFIG = "SELECT course_list FROM " + TEACHING_PLAN_TABLE + " WHERE faculty = ? AND speciality = ? AND academic_program = ? AND semester_order = ?;"

func (r *repository) GetListCourseOfTeachingPlan(faculty string, speciality string, academinProgram string, semester int) []dto.CourseInTeachingPlan {
	var listCourse  dto.CourseList
	rows, err := r.db.Queryx(SELECT_TEACHING_PLAN_CONFIG, faculty, speciality, academinProgram, semester)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		rows.Scan(&listCourse)
	}
	
	fmt.Printf("%+v", listCourse)
	
	
	
	return nil
}

