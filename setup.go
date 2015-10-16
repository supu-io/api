package main

import (
	"net/http"

	"github.com/go-martini/martini"
	"github.com/supu-io/messages"
)

// Setup the necessary tools to start using supu
func Setup(r *http.Request, params martini.Params) string {

	states, err := getStatuses()
	if err != nil {
		return "{\"error\":\"" + err.Error() + "\"}"
	}

	msg := messages.Setup{
		Org:    r.URL.Query().Get("org"),
		Repo:   r.URL.Query().Get("repo"),
		States: states,
		Config: config(),
	}

	return Request("issue-tracker.setup", msg)
}
