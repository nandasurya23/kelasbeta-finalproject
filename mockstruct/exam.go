package mockstruct

type RequestGetExam struct {
	Identifier    string         `json:"identifier"`
	Name          string         `json:"name"`
	ExamQuestions []ExamQuestion `json:"exam_questions"`
}

type ExamQuestion struct {
	ID            uint         `json:"id"`
	Question      string       `json:"question"`
	Order         uint         `json:"order"`
	CategoryName  string       `json:"category_name"`
	CategoryOrder uint         `json:"category_order"`
	Answers       []ExamAnswer `json:"answers"`
}

type ExamAnswer struct {
	ID      uint   `json:"id"`
	Opsi    string `json:"opsi"`
	Jawaban string `json:"jawaban"`
}
