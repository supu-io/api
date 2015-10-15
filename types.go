package main

import (
	"encoding/json"
	"log"
)

// Github configuration structure
type Github struct {
	Token string `json:"token"`
}

// Config structure
type Config struct {
	Github `json:"github"`
}

// IssuesList list of issues
type IssuesList struct {
	Status  string `json:"status"`
	Org     string `json:"org, omitempty"`
	Repo    string `json:"repo, omitempty"`
	*Config `json:"config"`
}

// Convert issues list to a json string
func (i *IssuesList) toJSON() []byte {
	json, err := json.Marshal(i)
	if err != nil {
		log.Println(err)
	}
	return json
}

// UpdateAttr json to update an issue
type UpdateAttr struct {
	Status string `json:"status"`
}
