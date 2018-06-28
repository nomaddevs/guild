package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/munsy/guild/models"
)

// Test page
func handleTest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		t := template.Must(template.ParseFiles(home + "/views/test.html"))
		t.Execute(w, nil)
		break
	default:
		fmt.Fprintln(w, "How did you get here...?")
	}
}
