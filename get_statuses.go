package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/go-martini/martini"
)

var source string

// Hook ..
type Hook struct {
	URL string `json:"url"`
}

// Transition ...
type Transition struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Hooks []Hook `json:"hooks"`
}

// Workflow ...
type Workflow struct {
	Transitions []Transition `json:"transitions"`
}

// GetStatuses get all statuses for the issue workflow
func GetStatuses(r *http.Request, params martini.Params) string {
	statuses, err := getStatuses()
	if err != nil {
		return GenerateErrorMessage(err)
	}
	res, err := json.Marshal(statuses)
	if err != nil {
		return GenerateErrorMessage(err)
	}

	return string(res)
}

// Get a list of valid statuses
func getStatuses() ([]string, error) {
	status := []string{}
	w := getWorkflow()
	for _, t := range w.Transitions {
		addFrom := true
		addTo := true
		for _, s := range status {
			if s == t.To {
				addTo = false
			}
			if s == t.From {
				addFrom = false
			}
		}
		if addFrom == true {
			status = append(status, string(t.From))
		}
		if addTo == true {
			status = append(status, string(t.To))
		}
	}

	return status, nil
}

func getWorkflow() Workflow {
	if source == "" {
		source = "workflows/default.json"
	}
	file, err := os.Open(source)
	if err != nil {
		log.Panic("error:", err)
	}
	decoder := json.NewDecoder(file)
	w := Workflow{}
	err = decoder.Decode(&w)
	if err != nil {
		log.Println("Workflow " + source + " not found")
		log.Panic("error:", err)
	}

	return w
}

// Checks if the given status is valid or not
func isValidStatus(status string) (bool, []string) {
	statuses, _ := getStatuses()

	for _, s := range statuses {
		if string(s) == status {
			return true, statuses
		}
	}

	return false, statuses
}
