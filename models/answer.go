package models

type Answer struct {
	ID         uint   `db:"id"`
	Option     string `db:"option"`
	Response   string `db:"response"`
	Score      int    `db:"score"`
	QuestionID uint   `db:"question_id"`
}
