package main

import (
	"time"

	"github.com/go-martini/martini"
)

// GetIssue is the /issue/:issue and gets am Issue details
func GetIssue(params martini.Params) string {
	issue := params["issue"]
	if owner, ok := params["owner"]; ok {
		issue = owner + "/" + params["repo"] + "/" + params["issue"]
	}

	return getIssueDetails(issue)

}

func getIssueDetails(issue string) string {
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
