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

func GetExamQuestions(mdlIdentifier string) (req mockstruct.RequestGetExam, err error) {
	seed := time.Now().UnixNano()
	randUtil := rand.New(rand.NewSource(seed))

	mdlData := models.Module{
		Identifier: mdlIdentifier,
	}

	err = mdlData.GetByIdentifier(config.Postgres.DB)
	if err != nil {
		return mockstruct.RequestGetExam{}, err
	}

	req.Name = mdlData.Name
	req.Identifier = mdlData.Identifier
	req.ExamQuestions = []mockstruct.ExamQuestion{}
	for _, k := range mdlData.QuestionIDS {
		qst := models.Question{
			Model: models.Model{
				ID: uint(k),
			},
		}
		qstData, err := qst.GetByID(config.Postgres.DB)
		if err == nil {
			randUtil.Shuffle(len(qstData.Answers), func(i, j int) {
				qstData.Answers[i], qstData.Answers[j] = qstData.Answers[j], qstData.Answers[i]
			})

			qstDataFormatted := mockstruct.ExamQuestion{
				ID:            qstData.ID,
				Question:      qstData.Question,
				CategoryName:  qstData.Category.Name,
				CategoryOrder: uint(qstData.Category.Order),
				Answers:       []mockstruct.ExamAnswer{},
			}
			for index, ps := range qstData.Answers {
				qstDataFormatted.Answers = append(qstDataFormatted.Answers, mockstruct.ExamAnswer{
					ID:      ps.ID,
					Opsi:    fmt.Sprintf("v%d", ('A' + index)),
					Jawaban: ps.Jawaban,
				})
			}
			req.ExamQuestions = append(req.ExamQuestions, qstDataFormatted)
		} else {
			return mockstruct.RequestGetExam{}, err
		}
	}

	categoryGroups := make(map[uint][]mockstruct.ExamQuestion)
	for _, eq := range req.ExamQuestions {
		categoryGroups[eq.CategoryOrder] = append(categoryGroups[eq.CategoryOrder], eq)
	}

	for _, group := range categoryGroups {
		randUtil.Shuffle(len(group), func(i, j int) {
			group[i], group[j] = group[j], group[i]
		})
	}

	var sortedCategories []uint
	for k := range categoryGroups {
		sortedCategories = append(sortedCategories, k)
	}

	sort.Slice(sortedCategories, func(i, j int) bool {
		return sortedCategories[i] < sortedCategories[j]
	})

	req.ExamQuestions = []mockstruct.ExamQuestion{}
	for _, catOrder := range sortedCategories {
		group := categoryGroups[catOrder]
		req.ExamQuestions = append(req.ExamQuestions, group...)
	}

	for index := range req.ExamQuestions {
		req.ExamQuestions[index].Order = uint(index + 1)
	}

	return
}
