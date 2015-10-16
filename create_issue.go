package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/supu-io/messages"
)

// CreateAttr json to create an issue
type CreateAttr struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Org   string `json:"org"`
	Repo  string `json:"repo"`
}

// CreateIssue is the POST /issue/:issue and updates an Issue status
func CreateIssue(r *http.Request, params martini.Params) string {
	var t CreateAttr
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&t)

	if err != nil {
		return GenerateErrorMessage(err)
	}

	msg := messages.CreateIssue{
		Issue: &messages.Issue{
			Title: t.Title,
			Body:  t.Body,
			Repo:  t.Repo,
			Org:   t.Org,
		},
		Config: config(),
	}

	return Request("issues.create", msg)
}
