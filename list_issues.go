package main

import (
	"net/http"
	"strings"

	"github.com/go-martini/martini"
	"github.com/supu-io/messages"
)

// GetIssues is the /issues and get a list of Issues by its status
func GetIssues(r *http.Request, params martini.Params) string {
	status := r.URL.Query().Get("status")
	org := r.URL.Query().Get("org")
	repo := r.URL.Query().Get("repo")

	if status != "" {
		if valid, statuses := isValidStatus(status); valid == false {
			st := strings.Join(statuses, ", ")
			return getError("Invalid status, valid statuses: " + st)
		}
	}

	msg := messages.GetIssuesList{
		Status: status,
		Repo:   repo,
		Org:    org,
		Config: config(),
	}

	response := Request("issues.list", msg)

	return response
}
