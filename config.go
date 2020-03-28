package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Database DataBaseConfig
	Telegram struct {
		Token string `json:"token"`
	} `json:"telegram"`
}

type DataBaseConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
	User     string `json:"user"`
	Password string `json:"password"`
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
