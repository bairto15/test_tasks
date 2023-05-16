package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Port     string `json:"port"`
	Host     string `json:"address"`
	Name     string `json:"name"`
	Password string `json:"password"`
	DBName   string `json:"db_name"`
	Schema   string `json:"shema"`
}

func Get() (Config, error) {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		return Config{}, err
	}

	var config Config

	json.Unmarshal(file, &config)

	return config, nil
}
