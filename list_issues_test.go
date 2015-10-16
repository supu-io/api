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

func listIssues(org string, repo string, status string) *httptest.ResponseRecorder {
	m := setupRouter()

	var jsonStr = []byte("")
	request, _ := http.NewRequest("GET", "/issues?org="+org+"&repo="+repo+"&status="+status, bytes.NewBuffer(jsonStr))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	response := httptest.NewRecorder()
	m.ServeHTTP(response, request)

	return response
}
func TestListIssues(t *testing.T) {
	setup()
	Convey("When I list the issues for a repo", t, func() {
		subscribe("workflow.states.all", `["doing","todo"]`)
		subscribe("issues.list", "")
		response := listIssues("org", "repo", "todo")
		res := messages.GetIssuesList{}
		json.Unmarshal([]byte(response.Body.String()), &res)
		Convey("Then it should return a valid message", func() {
			So(res.Org, ShouldEqual, "org")
			So(res.Repo, ShouldEqual, "repo")
			So(res.Status, ShouldEqual, "todo")
		})
	})
	Convey("When I list the issues for a repo", t, func() {
		subscribe("workflow.states.all", `["doing","todo"]`)
		subscribe("issues.list", "")
		Convey("And status does not exist", func() {
			response := listIssues("org", "repo", "foo")
			Convey("Then it should return a valid message", func() {
				So(response.Body.String(), ShouldEqual, "{\"error\":\"Invalid status, valid statuses: doing, todo\"}")
			})
		})
	})
}
