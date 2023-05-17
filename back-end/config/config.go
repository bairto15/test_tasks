package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	PortServer    string `json:"port_server"`
	PortServerGin string `json:"port_server_gin"`
	Host          string `json:"host"`
	Name          string `json:"name"`
	Password      string `json:"password"`
	DBName        string `json:"db_name"`
	DBPort        string `json:"db_port"`
	Schema        string `json:"shema"`
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
