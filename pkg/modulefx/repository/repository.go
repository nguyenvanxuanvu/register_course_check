package repository

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/nguyenvanxuanvu/register_course_check/pkg/common"
	"github.com/nguyenvanxuanvu/register_course_check/pkg/dto"
	"golang.org/x/exp/slices"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}



const MIN_MAX_CREDIT_TABLE = "min_max_credit"
const SELECT_MIN_CREDIT_CONFIG = "SELECT min_credit, max_credit, white_list FROM " + MIN_MAX_CREDIT_TABLE + " WHERE academic_program = ? AND semester = ?;"
func (r *repository) GetMinMaxCredit(studentId int, academinProgram string, semester int) (int,int,error) {
	
	var rawMinCredits, rawMaxCredits []int
	var rawWhiteLists []string
	rows, err := r.db.Queryx(SELECT_MIN_CREDIT_CONFIG, academinProgram, semester)
	if err != nil {
		log.Fatal(err)
	}
	idx := 0
	var x, y int
	var z string
	for rows.Next() {
		rows.Scan(&x, &y, &z)
		rawMinCredits = append(rawMinCredits, x)
		rawMaxCredits = append(rawMaxCredits, y)
		rawWhiteLists = append(rawWhiteLists, z)
		idx++
	}

	if idx == 0 {
		return -1,-1, errors.New(common.NOT_FOUND_MIN_MAX_CREDIT_CONFIG)
	}

	var numMinCredits, numMaxCredits int = rawMinCredits[0], rawMaxCredits[0]
	for idx, ele := range rawWhiteLists{
		if ele != ""{
			var listStudent []int
			err = json.Unmarshal([]byte(ele), &listStudent)
			if err != nil {
				return -1, -1, err
			}
			if slices.Contains(listStudent, studentId){
				return rawMinCredits[idx], rawMaxCredits[idx], nil
			}
			
		}
	}


	
	return numMinCredits, numMaxCredits, nil
}

const TEACHING_PLAN_TABLE = "teaching_plan"
const SELECT_TEACHING_PLAN_CONFIG = "SELECT course_list, free_credit_info FROM " + TEACHING_PLAN_TABLE + " WHERE faculty = ? AND speciality = ? AND academic_program = ? AND semester_order = ?;"

func (r *repository) GetListCourseOfTeachingPlan(faculty string, speciality string, academinProgram string, semester int)  ([]string, []dto.FreeCreditInfo, error) {
	var listCourse, freeCreditInfo  string
	rows, err := r.db.Queryx(SELECT_TEACHING_PLAN_CONFIG, faculty, speciality, academinProgram, semester)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		rows.Scan(&listCourse, &freeCreditInfo)
	}

	
	var rtFreeCreditInfo []dto.FreeCreditInfo
	var rtListCourse []string
	if listCourse != ""{
		err = json.Unmarshal([]byte(listCourse), &rtListCourse)
		if err != nil {
			return nil, nil, err
		}
	}
	if freeCreditInfo != ""{
		err = json.Unmarshal([]byte(freeCreditInfo), &rtFreeCreditInfo)
		if err != nil {
			return nil, nil, err
		}
	}
	

	
	return rtListCourse, rtFreeCreditInfo, nil
}

