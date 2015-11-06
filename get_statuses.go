package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-martini/martini"
)

var source string

// GetStatuses get all statuses for the issue workflow
func GetStatuses(r *http.Request, params martini.Params) string {
	statuses, err := getStatuses()
	if err != nil {
		return GenerateErrorMessage(err)
	}
	res, err := json.Marshal(statuses)
	if err != nil {
		return GenerateErrorMessage(err)
	}

	return string(res)
}

// Get a list of valid statuses
func getStatuses() ([]string, error) {
	status := []string{}
	w := getWorkflow()
	for _, t := range w.Transitions {
		addFrom := true
		addTo := true
		for _, s := range status {
			if s == t.To {
				addTo = false
			}
			if s == t.From {
				addFrom = false
			}
		}
		if addFrom == true {
			status = append(status, string(t.From))
		}
		if addTo == true {
			status = append(status, string(t.To))
		}
	}

	return status, nil
}

// Checks if the given status is valid or not
func isValidStatus(status string) (bool, []string) {
	statuses, _ := getStatuses()

	for _, s := range statuses {
		if string(s) == status {
			return true, statuses
		}
	}

	return false, statuses
}
