package main

import (
	"strconv"

	"github.com/go-martini/martini"
	"github.com/supu-io/messages"
)

// GetIssue is the /issue/:issue and gets am Issue details
func GetIssue(params martini.Params) string {
	number, err := strconv.Atoi(params["issue"])
	if err != nil {
		return GenerateErrorMessage(err)
	}

	msg := messages.GetIssue{
		Issue: &messages.Issue{
			ID:     params["issue"],
			Number: number,
			Org:    params["owner"],
			Repo:   params["repo"],
		},
		Config: config(),
	}
	response := Request("issues.details", msg)

	return response
}
