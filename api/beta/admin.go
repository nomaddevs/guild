package beta

import (
	"net/http"

	"github.com/munsy/guild/errors"
)

func (a *API) Admin(w http.ResponseWriter, r *http.Request) {
	e := &errors.Error{
		Message: "not implemented",
		Package: "api.beta",
		Type:    "API",
		Method:  "Admin",
	}
	a.Error(w, e)
}
