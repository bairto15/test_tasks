package service

import (
	"fmt"
	"test_puzzle/package/repository"
)

//Получить все варианты
func (s *Service) GetVariants() (variants []repository.Variant, err error) {
	variants, err = s.repository.GetVariants()

	return
}

//Получить задание
func (s *Service) Task(idUser, count, idVariant, idTest string) (interface{}, error) {
	if count == "" && idVariant == "" {
		s.Logger.Warn("He указаны данные")
		return nil, fmt.Errorf("He указаны данные")
	}

	if count == "4" {
		answers, err := s.repository.GetAnswers(idTest)
		if err != nil {
			s.Logger.Error(err)
			return nil, err
		}

		var countCorrAnswer int
		for _, v := range answers {
			if v.Answer == v.CorrAnswer {
				countCorrAnswer++
			}
		}

		s.mutex.Lock() 
		
		s.repository.AddResult(idTest, idUser, (countCorrAnswer*100/len(answers)*100)/100)
		
		s.mutex.Unlock()

		task, err := s.repository.GetResult(idTest)

		res := map[string]interface{}{"result": task}

		return res, err
	}

	task, err := s.repository.GetTasks(count, idVariant)

	if count == "1" {
		idTest, err := s.repository.StartTest(idUser, idVariant)

		res := map[string]interface{}{
			"task":   task,
			"idTest": idTest,
		}

		return res, err
	}

	res := map[string]interface{}{"task": task}

	return res, err
}

//Получить результат
func (s *Service) GetResult(idTest string) (result repository.Result, err error) {
	result, err = s.repository.GetResult(idTest)

	return
}

//Добавить ответ
func (s *Service) AddAnswer(idTest, idUser, answer, corrAnswer string) (err error) {
	if answer == "" && corrAnswer != "" {
		return fmt.Errorf("Не указан ответ")
	}

	if answer != "" {
		err = s.repository.AddAnswer(idTest, idUser, answer, corrAnswer)
	}

	return
}
