package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-martini/martini"
)

func GetStatuses(r *http.Request, params martini.Params) string {
	err, statuses := getStatuses()
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

func getStatuses() (error, []string) {
	body := "{\"issue\":{\"id\":\"\",\"status\":\"\"}}"
	res, err := nc.Request("workflow.states.all", []byte(body), 10000*time.Millisecond)
	if err != nil {
		return err, []string{}
	}

	statuses := []string{}
	err = json.Unmarshal(res.Data, &statuses)
	if err != nil {
		return err, []string{}
	}

	return err, statuses
}

// Checks if the given status is valid or not
func isValidStatus(status string) (bool, []string) {
	err, statuses := getStatuses()
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
