package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-martini/martini"
)

// UpdateIssueType ...
type UpdateIssueType struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

// UpdateIssueMsg ...
type UpdateIssueMsg struct {
	Issue   UpdateIssueType `json:"issue"`
	*Config `json:"config"`
	State   string `json:"state"`
}

func (i *UpdateIssueMsg) toJSON() []byte {
	json, err := json.Marshal(i)
	if err != nil {
		log.Println(err)
	}
	return json

}

// UpdateIssue is the PUT /issue/:issue and updates an Issue status
func UpdateIssue(r *http.Request, params martini.Params) string {
	decoder := json.NewDecoder(r.Body)
	var t UpdateAttr
	err := decoder.Decode(&t)
	status := t.Status
	if valid, statuses := isValidStatus(status); valid == false {
		st := strings.Join(statuses, ", ")
		return getError("Invalid status, valid statuses: " + st)
	}

	fullID := params["issue"]
	if params["owner"] != "" && params["repo"] != "" {
		fullID = params["owner"] + "/" + params["repo"] + "/" + params["issue"]
	}

	// TODO : Get issue current status
	data := getIssueDetails(fullID)
	type tmp struct {
		Status string `json:"status"`
	}
	details := tmp{}
	json.Unmarshal([]byte(data), &details)

	msg := UpdateIssueMsg{
		Issue: UpdateIssueType{
			ID:     fullID,
			Status: details.Status,
		},
		State:  status,
		Config: getConfig(),
	}
	issues, err := nc.Request("workflow.move", msg.toJSON(), 10000*time.Millisecond)
	if err != nil {
		return "{\"error\":\"" + err.Error() + "\"}"
	}
	log.Println(string(issues.Data))

	return string(issues.Data)
}
