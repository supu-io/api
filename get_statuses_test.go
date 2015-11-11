package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func getStatusesRequest() *httptest.ResponseRecorder {
	m := setupRouter()

	var jsonStr = []byte("")
	request, _ := http.NewRequest("GET", "/statuses", bytes.NewBuffer(jsonStr))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	response := httptest.NewRecorder()
	m.ServeHTTP(response, request)

	return response
}

func TestListAllStatuses(t *testing.T) {
	setup()
	Convey("When I list the issues for a repo", t, func() {
		source = "workflows/test/default.json"
		response := getStatusesRequest()
		Convey("Then it should return a valid message", func() {
			So(response.Body.String(), ShouldEqual, `["doing","todo"]`)
		})
	})
}
