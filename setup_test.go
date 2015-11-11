package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/supu-io/messages"

	. "github.com/smartystreets/goconvey/convey"
)

func issueTrackerSetup(org string, repo string) *httptest.ResponseRecorder {
	m := setupRouter()

	var jsonStr = []byte("")
	request, _ := http.NewRequest("POST", "/setup?org="+org+"&repo="+repo, bytes.NewBuffer(jsonStr))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	response := httptest.NewRecorder()
	m.ServeHTTP(response, request)

	return response
}
func TestSetup(t *testing.T) {
	setup()
	Convey("When I setup a repo", t, func() {
		source = "workflows/test/default.json"
		subscribe("issue-tracker.setup", "")
		response := issueTrackerSetup("org", "repo")
		res := messages.Setup{}
		json.Unmarshal([]byte(response.Body.String()), &res)
		Convey("Then it should return a valid message", func() {
			So(res.Org, ShouldEqual, "org")
			So(res.Repo, ShouldEqual, "repo")
			So(len(res.States), ShouldEqual, 2)
		})
	})
}
