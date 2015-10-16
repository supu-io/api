package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-martini/martini"
)

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
	res := Request("workflow.states.all", "")

	statuses := []string{}
	err := json.Unmarshal([]byte(res), &statuses)
	if err != nil {
		return []string{}, err
	}

	return statuses, err
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
