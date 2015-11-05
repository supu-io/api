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

func getIssueRequest(org string, repo string, issue string) *httptest.ResponseRecorder {
	m := setupRouter()

	var jsonStr = []byte("")
	request, _ := http.NewRequest("GET", "/issues/"+org+"/"+repo+"/"+issue, bytes.NewBuffer(jsonStr))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	response := httptest.NewRecorder()
	m.ServeHTTP(response, request)

	return response
}
func TestGetIssue(t *testing.T) {
	setup()
	Convey("When I setup a repo", t, func() {
		source = "workflows/test/default.json"
		subscribe("issues.details", "")
		response := getIssueRequest("org", "repo", "1")
		res := messages.GetIssue{}
		println(response.Body.String())
		json.Unmarshal([]byte(response.Body.String()), &res)
		Convey("Then it should return a valid message", func() {
			issue := *res.Issue
			So(issue.Org, ShouldEqual, "org")
			So(issue.Repo, ShouldEqual, "repo")
			So(issue.ID, ShouldEqual, "1")
		})
	})
}
