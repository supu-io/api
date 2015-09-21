package main

import (
	"github.com/go-martini/martini"
	"net/http"
	"time"
)

func GetHome(w http.ResponseWriter, r *http.Request) string {
	return "Hello world"
}

func GetIssues(params martini.Params) string {
	body := []byte("{\"status\":\"" + params["status"] + "\"}")
	issues, err := nc.Request("issues.list", body, 10*time.Millisecond)
	if err != nil {
		return "error"
	}

	return string(issues.Data)
}

func GetIssue(params martini.Params) string {
	body := []byte("{\"issue\":\"" + params["issue"] + "\"}")
	issues, err := nc.Request("issues.details", body, 10*time.Millisecond)
	if err != nil {
		return "error"
	}

	return string(issues.Data)
}

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

func getError(body string) string {
	return "{\"error\":\"" + body + "\"}"
}

func isValidStatus(status string) bool {
	if status == "todo" || status == "uat" || status == "doing" || status == "done" || status == "review" {
		return true
	}
	return false
}
