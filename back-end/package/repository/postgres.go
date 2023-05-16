package repository

import (
	"test_puzzle/package/logging"

	"github.com/jmoiron/sqlx"
)

type Postgres struct {
	logger logging.Logger
	db     *sqlx.DB
}

func NewPostgres(db *sqlx.DB) *Postgres {
	return &Postgres{
		db:     db,
		logger: logging.GetLogger(),
	}
}

//Пользователь
func (r *Postgres) GetUser(login, password string) (user User, err error) {
	query := "SELECT id FROM users WHERE login=$1 AND password=$2"
	
	err = r.db.Get(&user, query, login, password)
	if err != nil {
		r.logger.Error(err)
	}

	return
}

//Варианты ответов
func (r *Postgres) GetVariants() (variants []Variant, err error) {
	query := "SELECT * FROM variants"

	rows, err := r.db.Query(query)
	if err != nil {
		r.logger.Error(err)
	}
	defer rows.Close()

	for rows.Next(){
		variant := Variant{}

		err = rows.Scan(&variant.Id, &variant.Name)
		if err != nil{
            r.logger.Error(err)
            continue
        }
		
		variants = append(variants, variant)
	}

	return
}

//Вопросы, варианты ответов, правильный ответ, 
func (r *Postgres) GetTasks(count, idVariant string) (tasks Tasks, err error) {
	query := "SELECT * FROM tasks WHERE count=$1 AND variant_id=$2"

	err = r.db.Get(&tasks, query, count, idVariant)
	if err != nil {
		r.logger.Error(err)
	}

	return
}

//Началов теста
func (r *Postgres) StartTest(idUser, idVariant string) (idTest string, err error) {
	query := "INSERT INTO tests (user_id, variant_id, date) values ($1, $2, current_timestamp) RETURNING id"

	row := r.db.QueryRow(query, idUser, idVariant)
	if err = row.Scan(&idTest); err != nil {
		r.logger.Error(err)
	}

	return
}

//Получить ответ по id
func (r *Postgres) GetAnswer(idAnswer string) (answers Answers, err error) {
	query := "SELECT * FROM answers WHERE id=$1"

	err = r.db.Select(&answers, query, idAnswer)
	if err != nil {
		r.logger.Error(err)
	}

	return
}

//Добавить ответ
func (r *Postgres) AddAnswer(idTest, idUser, answer, corrAnswer string) (err error) {
	query := "INSERT INTO answers (test_id, user_id, answer, correct_answer) values ($1, $2, $3, $4)"

	_, err = r.db.Exec(query, idTest, idUser , answer, corrAnswer)
	if err != nil {
		r.logger.Error(err)
	}

	return
}

//Итог тестирования
func (r *Postgres) AddResult(idTest, idUser string, persent int) (err error) {
	query := "INSERT INTO results(test_id, user_id, percent) values ($1, $2, $3)"

	_, err = r.db.Exec(query, idTest, idUser, persent)
	if err != nil {
		r.logger.Error(err)
	}

	return
}

//Итог тестирования
func (r *Postgres) GetResult(idTest string) (result Result, err error) {
	query := "SELECT * FROM results WHERE test_id=$1"

	err = r.db.Get(&result, query, idTest)
	if err != nil {
		r.logger.Error(err)
	}

	return
}

//Варианты ответов
func (r *Postgres) GetAnswers(idTest string) (answers []Answers, err error) {
	query := "SELECT id, answer, correct_answer FROM answers WHERE test_id=$1"

	rows, err := r.db.Query(query, idTest)
	if err != nil {
		r.logger.Error(err)
	}
	defer rows.Close()

	for rows.Next(){
		ans := Answers{}

		err = rows.Scan(&ans.Id, &ans.Answer, &ans.CorrAnswer)
		if err != nil{
            r.logger.Error(err)
            continue
        }
		
		answers = append(answers, ans)
	}

	return
}

//Авторизация пользователя
func (r *Postgres) Auth(login string) (err error) {
	query := "UPDATE users SET auth=true, date_auth=current_timestamp WHERE login=$1"
	_, err = r.db.Exec(query, login)
	
	return
}

//Авторизация пользователя
func (r *Postgres) Out(login string) (err error) {
	query := "UPDATE users SET auth=false, date_out=current_timestamp WHERE login=$1"
	_, err = r.db.Exec(query, login)

	return
}
