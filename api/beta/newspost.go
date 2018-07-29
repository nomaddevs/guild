package beta

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/munsy/guild/models"
)

func handleMakeNewsPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Sorry, nothing here!")
}
