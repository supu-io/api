package main

import (
	"net/http"
)

// GetHome is the / route and gets a "Home page"
func GetHome(w http.ResponseWriter, r *http.Request) string {
	return "One single tool to rule them all"
}
