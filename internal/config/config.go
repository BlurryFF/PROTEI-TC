package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	GRPCConfig struct {
		IP         string `json:"ip"`
		Port       string `json:"port"`
		QueueSize  int    `json:"queueSize"`
		NumWorkers int    `json:"numWorkers"`
	} `json:"grpc"`
	HTTPConfig struct {
		IP   string `json:"ip"`
		Port string `json:"port"`
		Auth struct {
			Username string `json:"username"`
			Pass     string `json:"pass"`
		} `json:"auth"`
		AbsencesFileName  string `json:"absencesFileName"`
		EmployeesFileName string `json:"employeesFileName"`
		BaseStoragePath   string `json:"baseStoragePath"`
	} `json:"http"`
}

func LoadConfiguration(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	if err != nil {
		log.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}
