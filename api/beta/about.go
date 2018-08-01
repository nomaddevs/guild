package beta

import (
	"fmt"
	"net/http"
)

// About page
func (a *API) About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Sorry, nothing here!")
}
