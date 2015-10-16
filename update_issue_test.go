package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func updateIssueStatus(id string, status string) *httptest.ResponseRecorder {
	m := setupRouter()

	var jsonStr = []byte(`{"status":"` + status + `"}`)
	request, _ := http.NewRequest("PUT", "/issues/1", bytes.NewBuffer(jsonStr))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	response := httptest.NewRecorder()
	m.ServeHTTP(response, request)

	return response
}

func TestUpdateIssue(t *testing.T) {
	setup()

	Convey("When I update an issue with valid status", t, func() {
		subscribe("workflow.states.all", `["doing","todo"]`)
		subscribe("workflow.move", `update test`)
		go subscribe("issues.details", "update test")
		response := updateIssueStatus("1", "todo")
		Convey("Then it should return a valid message", func() {
			So(response.Body.String(), ShouldEqual, "update test")
		})
	})

	Convey("When I update an issue with an invalid status", t, func() {
		subscribe("workflow.states.all", `["doing","todo"]`)
		subscribe("workflow.move", `update test`)
		response := updateIssueStatus("1", "foo")
		Convey("Then it should return an error", func() {
			So(response.Body.String(), ShouldEqual, "{\"error\":\"Invalid status, valid statuses: doing, todo\"}")
		})
	})
}
