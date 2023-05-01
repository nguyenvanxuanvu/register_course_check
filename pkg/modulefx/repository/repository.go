package repository

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/nguyenvanxuanvu/register_course_check/pkg/common"
	"github.com/nguyenvanxuanvu/register_course_check/pkg/dto"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

const MIN_MAX_CREDIT_TABLE = "min_max_credit"

const WHITE_LIST_TABLE = "white_list"
const SELECT_MIN_MAX_CREDIT_CONFIG = "SELECT min_credit, max_credit FROM " + MIN_MAX_CREDIT_TABLE + " WHERE academic_program = ? AND semester = ?;"
const SELECT_MIN_MAX_CREDIT_FROM_WHITELIST = "SELECT min_credit, max_credit FROM " + WHITE_LIST_TABLE + " WHERE student_id = ? AND semester = ?;"

func (r *repository) GetMinMaxCredit(studentId string, academinProgram string, semester int) (int, int, error) {

	var minCredit, maxCredit int

	rows, err := r.db.Queryx(SELECT_MIN_MAX_CREDIT_FROM_WHITELIST, studentId, semester)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		rows.Scan(&minCredit, &maxCredit)
		return minCredit, maxCredit, nil
	}

	rows, err = r.db.Queryx(SELECT_MIN_MAX_CREDIT_CONFIG, academinProgram, semester)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		rows.Scan(&minCredit, &maxCredit)
		return minCredit, maxCredit, nil
	}

	return -1, -1, errors.New(common.NOT_FOUND_MIN_MAX_CREDIT_CONFIG)
}

const TEACHING_PLAN_TABLE = "teaching_plan"
const SELECT_TEACHING_PLAN_CONFIG = "SELECT course_list, free_credit_info FROM " + TEACHING_PLAN_TABLE + " WHERE faculty = ? AND speciality = ? AND academic_program = ? AND semester_order = ?;"

func (r *repository) GetListCourseOfTeachingPlan(faculty string, speciality string, academinProgram string, semester int) ([]string, []dto.FreeCreditInfo, error) {
	var listCourse, freeCreditInfo string
	rows, err := r.db.Queryx(SELECT_TEACHING_PLAN_CONFIG, faculty, speciality, academinProgram, semester)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		rows.Scan(&listCourse, &freeCreditInfo)
	}

	var rtFreeCreditInfo []dto.FreeCreditInfo
	var rtListCourse []string
	if listCourse != "" {
		err = json.Unmarshal([]byte(listCourse), &rtListCourse)
		if err != nil {
			return nil, nil, err
		}
	}
	if freeCreditInfo != "" {
		err = json.Unmarshal([]byte(freeCreditInfo), &rtFreeCreditInfo)
		if err != nil {
			return nil, nil, err
		}
	}

	return rtListCourse, rtFreeCreditInfo, nil
}


const UPDATE_COURSE_CONDITION = "UPDATE " + COURSE_CONDITION_TABLE + " SET course_condition = ? WHERE course_id = ?;"


func (r *repository) UpdateCourseCondition(listCourseCondition []dto.CourseConditionConfig) (bool, error) {
	for _, condition := range listCourseCondition{
		
		conditionJson, err := json.Marshal(condition.Condition)
		if err != nil {
			return false, err
		}
		_, err = r.db.Exec(UPDATE_COURSE_CONDITION, conditionJson , condition.CourseId)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}
