package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-martini/martini"
)

// GetHome is the / route and gets a "Home page"
func GetHome(w http.ResponseWriter, r *http.Request) string {
	return "One single tool to rule them all"
}

// GetIssues is the /issues and get a list of Issues by its status
// Status param is required
func GetIssues(r *http.Request, params martini.Params) string {
	status := r.URL.Query().Get("status")
	if valid, statuses := isValidStatus(status); valid == false {
		st := strings.Join(statuses, ", ")
		return getError("Invalid status, valid statuses: " + st)
	}
	config := getConfig()
	msg := IssuesList{
		Status: status,
		Org:    "supu-io",
		Config: config,
	}

	issues, err := nc.Request("issues.list", msg.toJSON(), 10000*time.Millisecond)
	if err != nil {
		return "{\"error\":\"" + err.Error() + "\"}"
	}

	return string(issues.Data)
}

// GetIssue is the /issue/:issue and gets am Issue details
func GetIssue(params martini.Params) string {
	issue := params["issue"]
	if owner, ok := params["owner"]; ok {
		issue = owner + "/" + params["repo"] + "/" + params["issue"]
	}

	config := getConfig()
	msg := IssueDetails{
		ID:     issue,
		Config: config,
	}
	issues, err := nc.Request("issues.details", msg.toJSON(), 10000*time.Millisecond)
	if err != nil {
		return "{\"error\":\"" + err.Error() + "\"}"
	}

	return string(issues.Data)
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

	body := []byte("{\"issue\":\"" + params["issue"] + "\", \"status\":\"" + status + "\"}")
	issues, err := nc.Request("workflow.move", body, 10000*time.Millisecond)
	if err != nil {
		return "{\"error\":\"" + err.Error() + "\"}"
	}

	return string(issues.Data)
}

// Returns a standard error message with the given body
func getError(body string) string {
	return "{\"error\":\"" + body + "\"}"
}

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
