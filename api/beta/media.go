package beta

import (
	"net/http"

	"github.com/munsy/guild/errors"
)

// Media page
func (a *API) Media(w http.ResponseWriter, r *http.Request) {
	e := &errors.Error{
		Message: "not implemented",
		Package: "api.beta",
		Type:    "API",
		Method:  "Media",
	}
	a.Error(w, e)
}
