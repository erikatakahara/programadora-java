package infra

import (
    "encoding/json"
    "os"
)

type Config struct {
	GcmPushKey string
}

func Conf() Config {
	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	config := Config{}
	decoder.Decode(&config)
	return config
}