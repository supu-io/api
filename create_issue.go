package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-martini/martini"
)

// CreateIssueType ...
type CreateIssueType struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Org   string `json:"owner"`
	Repo  string `json:"repo"`
}

// CreateIssueMsg ...
type CreateIssueMsg struct {
	Issue   CreateIssueType `json:"issue"`
	*Config `json:"config"`
}

// CreateAttr json to create an issue
type CreateAttr struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Org   string `json:"org"`
	Repo  string `json:"repo"`
}

func (i *CreateIssueMsg) toJSON() []byte {
	json, err := json.Marshal(i)
	if err != nil {
		log.Println(err)
	}
	return json

}

// CreateIssue is the POST /issue/:issue and updates an Issue status
func CreateIssue(r *http.Request, params martini.Params) string {
	var t CreateAttr

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&t)
	if err != nil {
		return "{\"error\":\"" + err.Error() + "\"}"
	}

	msg := CreateIssueMsg{
		Issue: CreateIssueType{
			Title: t.Title,
			Body:  t.Body,
			Repo:  t.Repo,
			Org:   t.Org,
		},
		Config: getConfig(),
	}

	issue, err := nc.Request("issues.create", msg.toJSON(), 10000*time.Millisecond)
	if err != nil {
		return "{\"error\":\"" + err.Error() + "\"}"
	}

	return string(issue.Data)
}
