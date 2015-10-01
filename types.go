package main

import (
	"encoding/json"
	"log"
)

type Github struct {
	Token string `json:"token"`
}

type Config struct {
	Github `json:"github"`
}

type IssuesList struct {
	Status  string `json:"status"`
	Org     string `json:"org, omitempty"`
	Repo    string `json:"repo, omitempty"`
	*Config `json:"config"`
}

func (i *IssuesList) toJSON() []byte {
	json, err := json.Marshal(i)
	if err != nil {
		log.Println(err)
	}
	return json
}

type IssueDetails struct {
	ID      string `json:"id"`
	*Config `json:"config"`
}

func (i *IssueDetails) toJSON() []byte {
	json, err := json.Marshal(i)
	if err != nil {
		log.Println(err)
	}
	return json
}

type UpdateAttr struct {
	Status string `json:"status"`
}
