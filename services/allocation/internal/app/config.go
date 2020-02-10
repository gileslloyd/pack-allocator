package app

import (
	"encoding/json"
	"os"
)

type Rabbit struct {
	host     string
	port     int
	user     string
	password string
	queue    string
}

func GetRabbitConfig() Rabbit {
	file, _ := os.Open("config/rabbit.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Rabbit{}
	err := decoder.Decode(&configuration)

	if err != nil {
		panic("Could not load Rabbit config")
	}

	return configuration
}
