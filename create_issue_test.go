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

func createIssueRequest(org string, repo string, title string, body string) *httptest.ResponseRecorder {
	m := setupRouter()
	var jsonStr = []byte(`{"title":"` + title + `","body":"` + body + `","repo":"` + repo + `","org":"` + org + `"}`)
	request, _ := http.NewRequest("POST", "/issues", bytes.NewBuffer(jsonStr))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	response := httptest.NewRecorder()
	m.ServeHTTP(response, request)

	return response
}

func TestCreateIssue(t *testing.T) {
	setup()
	Convey("When I create an issue", t, func() {
		subscribe("issues.create", "")
		response := createIssueRequest("org", "repo", "title", "body")
		res := messages.CreateIssue{}
		println(response.Body.String())
		json.Unmarshal([]byte(response.Body.String()), &res)
		Convey("Then it should return a valid message", func() {
			issue := *res.Issue
			So(issue.Org, ShouldEqual, "org")
			So(issue.Repo, ShouldEqual, "repo")
			So(issue.Title, ShouldEqual, "title")
			So(issue.Body, ShouldEqual, "body")
		})
	})
}
