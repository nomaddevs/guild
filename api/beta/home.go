package beta

import (
	"fmt"
	"net/http"
)

// Index page
func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Sorry, nothing here!")
}
