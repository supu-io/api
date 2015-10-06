package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-martini/martini"
)

// UpdateIssue is the PUT /issue/:issue and updates an Issue status
func UpdateIssue(r *http.Request, params martini.Params) string {
	decoder := json.NewDecoder(r.Body)
	var t UpdateAttr
	err := decoder.Decode(&t)
	status := t.Status
	if valid, statuses := isValidStatus(status); valid == false {
		st := strings.Join(statuses, ", ")
		return getError("Invalid status, valid statuses: " + st)
	}
	log.Println("....x.")

	// TODO: Hardcoded!!!!
	originalStatus := "todo"
	body := []byte(`{"issue":{"id":"` + params["issue"] + `", "status":"` + originalStatus + `"},"state":"` + status + `"}`)
	log.Println("....y.")
	issues, err := nc.Request("workflow.move", body, 10000*time.Millisecond)
	log.Println("....y.")
	if err != nil {
		log.Println(".....z")
		log.Println(body)
		return "{\"error\":\"" + err.Error() + "\"}"
	}
	log.Println(string(issues.Data))

	return string(issues.Data)
}
