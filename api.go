package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/supu-io/messages"
)

// GetHome is the / route and gets a "Home page"
func GetHome(w http.ResponseWriter, r *http.Request) string {
	return "One single tool to rule them all"
}

// Returns a standard error message with the given body
// TODO : DEPRECATED -> GenerateErrorMessage
func getError(body string) string {
	return "{\"error\":\"" + body + "\"}"
}

func config() messages.Config {
	c := messages.Config{}
	file, err := os.Open("config.json")
	if err != nil {
		log.Panic("error:", err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&c)
	if err != nil {
		log.Println("Config file is invalid")
		log.Panic("error:", err)
	}
	return c
}

// TODO : DEPRECATED -> config()
func getConfig() *Config {
	c := Config{}
	file, err := os.Open("config.json")
	if err != nil {
		log.Panic("error:", err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&c)
	if err != nil {
		log.Println("Config file is invalid")
		log.Panic("error:", err)
	}
	return &c
}
