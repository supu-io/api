package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/go-martini/martini"
)

// GetIssues is the /issues and get a list of Issues by its status
// Status param is required
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
	config := getConfig()
	msg := IssuesList{
		Status: status,
		Repo:   repo,
		Org:    org,
		Config: config,
	}

	issues, err := nc.Request("issues.list", msg.toJSON(), 10000*time.Millisecond)
	if err != nil {
		return "{\"error\":\"" + err.Error() + "\"}"
	}

	return string(issues.Data)
}
