package beta

import (
	"net/http"
	//"github.com/munsy/battlenet"
	"github.com/munsy/guild/errors"
)

// Recruitment handles recruitment data
func (a *API) Recruitment(w http.ResponseWriter, r *http.Request) {
	e := &errors.Error{
		Message: "not implemented",
		Package: "api.beta",
		Type:    "API",
		Method:  "Recruitment",
	}
	a.Error(w, e)
}
