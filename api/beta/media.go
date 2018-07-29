package beta

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/munsy/guild/models"
)

// Media page
func handleMedia(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Sorry, nothing here!")
}
