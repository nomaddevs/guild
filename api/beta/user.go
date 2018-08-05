package beta

import (
	"net/http"

	"github.com/munsy/battlenet"
	"github.com/munsy/guild/errors"
)

// User handles data for the logged in user.
func (a *API) User(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		r.ParseForm()

		token, err := r.Cookie("token")

		if nil != err {
			e := &errors.Error{
				Message: err.Error(),
				Package: "api.beta",
				Type:    "API",
				Method:  "User",
			}
			a.Error(w, e)
			return
		}

		client, err := battlenet.AccountClient(a.settings, token.Value)

		if nil != err {
			e := &errors.Error{
				Message: err.Error(),
				Package: "api.beta",
				Type:    "API",
				Method:  "User",
			}
			a.Error(w, e)
			return
		}

		bid, err := client.BattleID()

		if nil != err {
			e := &errors.Error{
				Message: err.Error(),
				Package: "api.beta",
				Type:    "API",
				Method:  "User",
			}
			a.Error(w, e)
			return
		}

		u := &models.User{
			ID:        bid.ID,
			BattleTag: bid.BattleTag,
			Applied:   models.Applied(bid.ID),
		}

		a.JSON(w, u)
		break
	default:
		e := &errors.Error{
			Message: "default hit",
			Package: "api.beta",
			Type:    "API",
			Method:  "User",
		}
		a.Error(w, e)
		return
	}
}
