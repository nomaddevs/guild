package beta

import (
	"net/http"
	//"github.com/munsy/battlenet"
	"github.com/munsy/guild/errors"
)

// Progression handles guild progression data
func (a *API) Progression(w http.ResponseWriter, r *http.Request) {
	e := &errors.Error{
		Message: "not implemented",
		Package: "api.beta",
		Type:    "API",
		Method:  "Progression",
	}
	a.Error(w, e)
}
