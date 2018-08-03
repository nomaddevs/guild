package beta

import (
	"net/http"

	//"github.com/munsy/battlenet"
	"github.com/munsy/guild/errors"
)

// MythicPlus handles Mythic+ data
func (a *API) MythicPlus(w http.ResponseWriter, r *http.Request) {
	e := &errors.Error{
		Message: "not implemented",
		Package: "api.beta",
		Type:    "API",
		Method:  "MythicPlus",
	}
	a.Error(w, e)
}
