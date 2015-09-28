package main

import (
	"encoding/json"
	"github.com/go-martini/martini"
	"log"
	"net/http"
	"os"
	"time"
)

// Get a "Home page"
func GetHome(w http.ResponseWriter, r *http.Request) string {
	return "One single tool to rule them all"
}

// Get a list of issues by its status
// Status param is required
func GetIssues(r *http.Request, params martini.Params) string {
	status := r.URL.Query().Get("status")
	if false == isValidStatus(status) {
		return getError("Invalid status")
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

// Get am Issue details
func GetIssue(params martini.Params) string {
	body := []byte("{\"issue\":\"" + params["issue"] + "\"}")
	issues, err := nc.Request("issues.details", body, 10*time.Millisecond)
	if err != nil {
		return "{\"error\":\"" + err.Error() + "\"}"
	}

	return string(issues.Data)
}

// Update Issue status
func UpdateIssue(r *http.Request, params martini.Params) string {
	status := r.FormValue("status")
	if false == isValidStatus(status) {
		return getError("Invalid status")
	}

	body := []byte("{\"issue\":\"" + params["issue"] + "\", \"status\":\"" + status + "\"}")
	issues, err := nc.Request("issues.details", body, 10*time.Millisecond)
	if err != nil {
		return "{\"error\":\"" + err.Error() + "\"}"
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
