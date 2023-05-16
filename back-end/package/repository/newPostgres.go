package repository

import (
	"fmt"
	"io/ioutil"
	"test_puzzle/config"
	"test_puzzle/package/logging"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresDB(conf config.Config) *sqlx.DB {
	logger := logging.GetLogger()

	cr := "user=postgres password=qwerty host=localhost port=5432 dbname=postgres sslmode=disable"

	db, err := sqlx.Connect("postgres", cr)
	if err != nil {
		logger.Fatal(err)
	}

	db.Exec("DROP SCHEMA tasks CASCADE")
	
	_, err = db.Exec(fmt.Sprintf(
		"CREATE SCHEMA IF NOT EXISTS %s; SET search_path TO %s", 
		conf.Schema, conf.Schema,
	))
	if err != nil {
		logger.Fatal(err)
	}
	
	sqlFile, err := ioutil.ReadFile("postgres.sql")
	if err != nil {
		logger.Fatal(err)
	}

	_, err = db.Exec(string(sqlFile))
	if err != nil {
		logger.Fatal(err)
	}

	return db
}
