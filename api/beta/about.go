package beta

import (
	"net/http"

	"github.com/munsy/guild/errors"
)

// About page
func (a *API) About(w http.ResponseWriter, r *http.Request) {
	e := &errors.Error{
		Message: "not implemented",
		Package: "api.beta",
		Type:    "API",
		Method:  "About",
	}
	a.Error(w, e)
}
