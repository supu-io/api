package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

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
