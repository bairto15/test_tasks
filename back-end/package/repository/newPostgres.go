package repository

import (
	"fmt"
	"io/ioutil"
	"test_puzzle/config"
	"test_puzzle/package/logging"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//docker run --name=postgres -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 -d --rm postgres

func NewPostgresDB(conf config.Config) *sqlx.DB {
	logger := logging.GetLogger()

	cr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		conf.Name, conf.Password, conf.Host, conf.DBPort, conf.DBName)

	logger.Info(cr)

	db, err := sqlx.Connect("postgres", cr)
	if err != nil {
		logger.Fatal(err)
	}

	db.Exec(fmt.Sprintf("DROP SCHEMA %s CASCADE", conf.Schema))

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
