package beta

import (
	"fmt"
	"net/http"
)

// Media page
func handleMedia(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Sorry, nothing here!")
}
