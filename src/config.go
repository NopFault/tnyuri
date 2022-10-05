package tnyuri

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Domain     string `json: "domain"`
	Database   string `json: "database"`
	Maxperuser int    `json: "maxperuser"`
	Token      string `json: "token"`
}

func GetConfig() Config {
	p, err := os.Getwd()
	if err != nil {
		log.Fatal("Cant get current working directory...")
	}
	content, err := ioutil.ReadFile(p + "/config.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var config Config
	err = json.Unmarshal(content, &config)
	if err != nil {
		log.Fatal("Error parsing config.json: ", err)
	}

	config.Database = p + config.Database

	return config
}
