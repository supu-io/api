package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/supu-io/messages"
)

// ToJSON represents an object as a json []byte
func ToJSON(i interface{}) []byte {
	json, err := json.Marshal(i)
	if err != nil {
		log.Println(err)
	}

	return json
}

// GenerateErrorMessage : Generates an error message from an error
// sent through an api response
func GenerateErrorMessage(err error) string {
	return GetErrorMessage(err.Error())
}

// GetErrorMessage : Generates an error message from a string
// sent through an api response
func GetErrorMessage(err string) string {
	return "{\"error\":\"" + err + "\"}"
}

// Request : does a nats request over the given subject and
// struct
func Request(subject string, msg interface{}) string {
	wait := 10000 * time.Millisecond
	body := ToJSON(msg)
	res, err := nc.Request(subject, body, wait)

	if err != nil {
		return GenerateErrorMessage(err)
	}

	return string(res.Data)
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

func getWorkflow() messages.Workflow {
	if source == "" {
		source = "workflows/default.json"
	}
	file, err := os.Open(source)
	if err != nil {
		log.Panic("error:", err)
	}
	decoder := json.NewDecoder(file)
	w := messages.Workflow{}
	err = decoder.Decode(&w)
	if err != nil {
		log.Println("Workflow " + source + " not found")
		log.Panic("error:", err)
	}

	return w
}
