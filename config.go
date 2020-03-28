package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Database DataBaseConfig
}

type DataBaseConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port": 3306,`
	Database string `json:"database": "bot_1"`
	User     string `json:"user": "bot1"`
	Password string `json:"password": "bot1"`
}

func loadConf(path string) Config {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	config := Config{}
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}
