package repository

type Store interface {
	GetUser(string, string) (User, error)
	GetVariants() ([]Variant, error)
	GetTasks(string, string) (Tasks, error)
	StartTest(string, string) (string, error)
	GetAnswer(string) (Answers, error)
	GetAnswers(string) ([]Answers, error)
	AddAnswer(string, string, string, string) error
	AddResult(string, string, int) error
	GetResult(string) (Result, error)
	Auth(string) error
	Out(string) error
}

type User struct {
	Id       string `db:"id"`
	Login    string `db:"login"`
	Password string `db:"password"`
	Auth     bool   `db:"auth"`
	DateAuth string `db:"date_auth"`
	DateOut  string `db:"date_out"`
}

type Variant struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}

type Tasks struct {
	Id            int    `db:"id"`
	IdVariand     int    `db:"variant_id"`
	Task          string `db:"task"`
	Count         int    `db:"count"`
	CorrectAnswer string `db:"correct_answer"`
	Answer1       string `db:"answer_1"`
	Answer2       string `db:"answer_2"`
	Answer3       string `db:"answer_3"`
	Answer4       string `db:"answer_4"`
}

type Test struct {
	Id        int    `db:"id"`
	IdUser    int    `db:"user_id"`
	IdVariand int    `db:"variant_id"`
	Date      string `db:"date"`
}

type Answers struct {
	Id         int    `db:"id"`
	IdTest     int    `db:"test_id"`
	IdUser     string `db:"user_id"`
	Answer     string `db:"answer"`
	CorrAnswer string `db:"correct_answer"`
}

type Result struct {
	Id      int `db:"id"`
	IdTest  int `db:"test_id"`
	IdUser  int `db:"user_id"`
	Percent int `db:"percent"`
}
