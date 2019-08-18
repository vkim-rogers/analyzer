package models

type QuizAnswers struct {
	PersonId string `json:"personId"`
	Name string `json:"name"`
	Questions []quizAnswersQA `json:"questions"`

}

type quizAnswersQA struct {
	Question string	`json:"question"`
	Answer string `json:"answer"`
}

func (qa *QuizAnswers) GetMock() QuizAnswers{
	return QuizAnswers{
		PersonId: "1",
		Name: "Оценка качества жизни",
		Questions: []quizAnswersQA{
			{Question:"Гибкий, приспособленный", Answer:"3"},
			{Question:"Добросердечный", Answer:"2"},
			{Question:"Открытый", Answer:"1"},
		},
	}
}