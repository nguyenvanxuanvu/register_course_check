package dto



type SubjectConfig struct {
	Id        string    `db:"id"`
	SubjectName string    `db:"subject_name"`
	NumCredits  int    `db:"num_credits"`
	Faculty string `db:"faculty"`
	

	SubjectConditionConfig []*SubjectCondtionConfig
}


type SubjectCondtionConfig struct {
	SubjectId     string `db:"subject_id"`
	SubjectDesId       string                     `db:"subject_des_id"`
	ConditionType int      `db:"condition_type"`   //1 : TQ    2: HT   3: SH 
}



