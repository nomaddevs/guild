package beta

import (
	"fmt"
	"net/http"
)

// About page
func handleAbout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Sorry, nothing here!")
}
