package beta

import (
	"fmt"
	"net/http"
)

// Media page
func (a *API) Media(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Sorry, nothing here!")
}
