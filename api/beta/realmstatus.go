package beta

import (
	"net/http"

	"github.com/munsy/battlenet"
	"github.com/munsy/guild/errors"
)

// RealmStatus handles realm status data
func (a *API) RealmStatus(w http.ResponseWriter, r *http.Request) {
	client, err := battlenet.WoWClient(a.settings, a.key)

	if nil != err {
		e := &errors.Error{
			Message: err.Error(),
			Package: "api.beta",
			Type:    "API",
			Method:  "RealmStatus",
		}
		a.Error(w, e)
		return
	}

	response, err := client.RealmStatus()

	if nil != err {
		e := &errors.Error{
			Message: err.Error(),
			Package: "api.beta",
			Type:    "API",
			Method:  "RealmStatus",
		}
		a.Error(w, e)
		return
	}

	switch r.Method {
	case "GET":
		a.JSON(w, response.Data)
		break
	default:
		e := &errors.Error{
			Message: "default hit",
			Package: "api.beta",
			Type:    "API",
			Method:  "RealmStatus",
		}
		a.Error(w, e)
		return
	}
}
