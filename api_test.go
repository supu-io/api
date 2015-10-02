package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nats-io/nats"

	. "github.com/smartystreets/goconvey/convey"
)

var w http.ResponseWriter
var r http.Request

func setup() {
	nc, _ = nats.Connect(nats.DefaultURL)
	c, _ = nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	// defer c.Close()
}

func subscribe(subject string, respond string) {
	sub, _ := nc.Subscribe(subject, func(m *nats.Msg) {
		println(respond)
		nc.Publish(m.Reply, []byte(respond))
	})
	sub.AutoUnsubscribe(1)
}

func TestGetHome(t *testing.T) {
	m := setupRouter()
	Convey("When I call GetHome", t, func() {

		request, _ := http.NewRequest("GET", "/", nil)
		response := httptest.NewRecorder()
		m.ServeHTTP(response, request)
		res := GetHome(w, &r)
		Convey("Then it should return a valid message", func() {
			So(res, ShouldEqual, "One single tool to rule them all")
		})
	})
}

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
		// go subscribe("issues.details", "update test")
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
