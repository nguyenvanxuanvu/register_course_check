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

