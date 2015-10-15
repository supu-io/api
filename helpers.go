package main

import (
	"encoding/json"
	"log"
	"time"
)

// ToJSON represents an object as a json []byte
func ToJSON(i interface{}) []byte {
	json, err := json.Marshal(i)
	if err != nil {
		log.Println(err)
	}

	return json
}

// GenerateErrorMessage : Generates an error message to be
// sent through an api response
func GenerateErrorMessage(err error) string {
	return "{\"error\":\"" + err.Error() + "\"}"
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
