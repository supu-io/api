package main

import (
	"net/http"
	"time"

	"github.com/go-martini/martini"
)

// GetHome is the / route and gets a "Home page"
func GetHome(w http.ResponseWriter, r *http.Request) string {
	return "One single tool to rule them all"
}

// GetIssues is the /issues and get a list of Issues by its status
// Status param is required
func GetIssues(params martini.Params) string {
	status := params["status"]
	if false == isValidStatus(status) {
		return getError("Invalid status")
	}
	body := []byte("{\"status\":\"" + status + "\"}")
	issues, err := nc.Request("issues.list", body, 10*time.Millisecond)
	if err != nil {
		return "error"
	}

	return string(issues.Data)
}

// GetIssue is the /issue/:issue and gets am Issue details
func GetIssue(params martini.Params) string {
	body := []byte("{\"issue\":\"" + params["issue"] + "\"}")
	issues, err := nc.Request("issues.details", body, 10*time.Millisecond)
	if err != nil {
		return "error"
	}

	return string(issues.Data)
}

// UpdateIssue is the PUT /issue/:issue and updates an Issue status
func UpdateIssue(r *http.Request, params martini.Params) string {
	status := r.FormValue("status")
	if false == isValidStatus(status) {
		return getError("Invalid status")
	}

	body := []byte("{\"issue\":\"" + params["issue"] + "\", \"status\":\"" + status + "\"}")
	issues, err := nc.Request("issues.details", body, 10*time.Millisecond)
	if err != nil {
		return "error: " + err.Error()
	}

	return string(issues.Data)
}

// Returns a standard error message with the given body
func getError(body string) string {
	return "{\"error\":\"" + body + "\"}"
}

// Checks if the given status is valid or not
func isValidStatus(status string) bool {
	if status == "todo" || status == "uat" || status == "doing" || status == "done" || status == "review" {
		return true
	}
	return false
}
