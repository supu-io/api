package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-martini/martini"
)

// Setup the necessary tools to start using supu
func Setup(r *http.Request, params martini.Params) string {

	type SetupMsg struct {
		Org     string   `json:"org"`
		Repo    string   `json:"repo"`
		States  []string `json:"states"`
		*Config `json:"config"`
	}

	states, err := getStatuses()
	if err != nil {
		return "{\"error\":\"" + err.Error() + "\"}"
	}

	msg := SetupMsg{
		Org:    r.URL.Query().Get("org"),
		Repo:   r.URL.Query().Get("repo"),
		States: states,
		Config: getConfig(),
	}

	body, err := json.Marshal(msg)
	if err != nil {
		return "{\"error\":\"" + err.Error() + "\"}"
	}

	issues, err := nc.Request("issue-tracker.setup", body, 10000*time.Millisecond)
	if err != nil {
		return "{\"error\":\"" + err.Error() + "\"}"
	}

	return string(issues.Data)
}
