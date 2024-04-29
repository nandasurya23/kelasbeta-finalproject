package utils

import (
	"fmt"
	"kelasbeta/finalproject/config"
	"kelasbeta/finalproject/mockstruct"
	"kelasbeta/finalproject/models"
	"math/rand"
	"sort"
	"time"
)

// GetExamQuestions mengambil pertanyaan ujian berdasarkan identifier modul.
func GetExamQuestions(mdlIdentifier string) (req mockstruct.RequestGetExam, err error) {
	// Seed random number generator
	seed := time.Now().UnixNano()
	randUtil := rand.New(rand.NewSource(seed))

	// Dapatkan data modul berdasarkan identifier
	mdlData := models.Module{
		Identifier: mdlIdentifier,
	}
	err = mdlData.GetByIdentifier(config.Postgres.DB)
	if err != nil {
		// Jika terjadi kesalahan saat mendapatkan modul, kembalikan error
		return mockstruct.RequestGetExam{}, err
	}

	// Atur nama dan identifier modul pada request
	req.Name = mdlData.Name
	req.Identifier = mdlData.Identifier
	req.ExamQuestions = []mockstruct.ExamQuestion{}

	// Iterasi melalui ID pertanyaan dalam modul
	for _, k := range mdlData.QuestionIDS {
		// Buat objek pertanyaan
		qst := models.Question{
			Model: models.Model{
				ID: uint(k),
			},
		}
		// Dapatkan data pertanyaan berdasarkan ID
		err := qst.GetByID(config.Postgres.DB)
		if err != nil {
			// Jika terjadi kesalahan saat mendapatkan pertanyaan, kembalikan error
			return mockstruct.RequestGetExam{}, err
		}

		// Kemas data pertanyaan ke dalam format yang diinginkan
		randUtil.Shuffle(len(qst.Answers), func(i, j int) {
			qst.Answers[i], qst.Answers[j] = qst.Answers[j], qst.Answers[i]
		})

		qstDataFormatted := mockstruct.ExamQuestion{
			ID:            qst.ID,
			Question:      qst.Question,
			CategoryName:  qst.Category.Name,
			CategoryOrder: uint(qst.Category.Order),
			Answers:       []mockstruct.ExamAnswer{},
		}
		for index, ps := range qst.Answers {
			qstDataFormatted.Answers = append(qstDataFormatted.Answers, mockstruct.ExamAnswer{
				ID:      ps.ID,
				Opsi:    fmt.Sprintf("v%d", ('A' + index)),
				Jawaban: ps.Jawaban,
			})
		}
		// Tambahkan pertanyaan yang telah diformat ke dalam request
		req.ExamQuestions = append(req.ExamQuestions, qstDataFormatted)
	}

	// Kelompokkan pertanyaan berdasarkan urutan kategori
	categoryGroups := make(map[uint][]mockstruct.ExamQuestion)
	for _, eq := range req.ExamQuestions {
		categoryGroups[eq.CategoryOrder] = append(categoryGroups[eq.CategoryOrder], eq)
	}

	// Acak setiap kelompok pertanyaan
	for _, group := range categoryGroups {
		randUtil.Shuffle(len(group), func(i, j int) {
			group[i], group[j] = group[j], group[i]
		})
	}

	// Urutkan kategori secara ascending
	var sortedCategories []uint
	for k := range categoryGroups {
		sortedCategories = append(sortedCategories, k)
	}
	sort.Slice(sortedCategories, func(i, j int) bool {
		return sortedCategories[i] < sortedCategories[j]
	})

	// Gabungkan pertanyaan dari setiap kategori berdasarkan urutan kategori
	req.ExamQuestions = []mockstruct.ExamQuestion{}
	for _, catOrder := range sortedCategories {
		group := categoryGroups[catOrder]
		req.ExamQuestions = append(req.ExamQuestions, group...)
	}

	// Atur nomor urut untuk setiap pertanyaan
	for index := range req.ExamQuestions {
		req.ExamQuestions[index].Order = uint(index + 1)
	}

	return
}
