package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-martini/martini"
)

// GetStatuses get all statuses for the issue workflow
func GetStatuses(r *http.Request, params martini.Params) string {
	statuses, err := getStatuses()
	if err != nil {
		return "{\"error\":\"Internal error\"}"
	}
	json, err := json.Marshal(statuses)
	if err != nil {
		log.Println(err)
		return "{\"error\":\"Internal error\"}"
	}

	return string(json)
}

func getStatuses() ([]string, error) {
	body := "{\"issue\":{\"id\":\"\",\"status\":\"\"}}"
	res, err := nc.Request("workflow.states.all", []byte(body), 10000*time.Millisecond)
	if err != nil {
		return []string{}, err
	}

	statuses := []string{}
	err = json.Unmarshal(res.Data, &statuses)
	if err != nil {
		return []string{}, err
	}

	return statuses, err
}

// Checks if the given status is valid or not
func isValidStatus(status string) (bool, []string) {
	statuses, err := getStatuses()
	if err != nil {
		return false, statuses
	}
	sw := false
	for _, s := range statuses {
		if string(s) == status {
			sw = true
		}
	}
	return sw, statuses
}
