package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-martini/martini"
	"github.com/supu-io/messages"
)

// UpdateAttr json to update an issue
type UpdateAttr struct {
	Status string `json:"status"`
}

// UpdateIssue is the PUT /issue/:issue and updates an Issue status
func UpdateIssue(r *http.Request, params martini.Params) string {
	decoder := json.NewDecoder(r.Body)
	var t UpdateAttr
	decoder.Decode(&t)
	status := t.Status
	if valid, statuses := isValidStatus(status); valid == false {
		st := strings.Join(statuses, ", ")
		return GetErrorMessage("Invalid status, valid statuses: " + st)
	}

	fullID := params["issue"]
	if params["owner"] != "" && params["repo"] != "" {
		fullID = params["owner"] + "/" + params["repo"] + "/" + params["issue"]
	}

	number, _ := strconv.Atoi(params["issue"])

	msg := messages.UpdateIssue{
		Issue: &messages.Issue{
			ID:     fullID,
			Number: number,
			Repo:   params["repo"],
			Org:    params["owner"],
			Status: issueCurrentStatus(params),
		},
		Status: status,
		Config: config(),
	}

	return Request("workflow.move", msg)
}

// Get the issue current status
func issueCurrentStatus(params martini.Params) string {
	data := GetIssue(params)
	type tmp struct {
		Status string `json:"status"`
	}
	details := tmp{}
	json.Unmarshal([]byte(data), &details)

	return details.Status
}
