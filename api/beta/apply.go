package beta

import (
	"net/http"

	"github.com/munsy/battlenet"
	"github.com/munsy/guild/errors"
	"github.com/munsy/guild/pkg/models"
)

func (a *API) Apply(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Get user token
		// If no token, show blizz auth
		c, err := r.Cookie("token")

		if nil != err {
			e := &errors.Error{
				Message: err.Error(),
				Package: "api.beta",
				Type:    "API",
				Method:  "Apply",
			}
			a.Error(w, e)
			return
		}

		client, err := battlenet.AccountClient(a.settings, c.Value)

		if nil != err {
			e := &errors.Error{
				Message: err.Error(),
				Package: "api.beta",
				Type:    "API",
				Method:  "Apply",
			}
			a.Error(w, e)
			return
		}

		// Send character data for them to choose from for app process
		response, err := client.WoWOauthProfile()

		if nil != err {
			e := &errors.Error{
				Message: err.Error(),
				Package: "api.beta",
				Type:    "API",
				Method:  "Apply",
			}
			a.Error(w, e)
			return
		}

		a.JSON(w, response.Data)
		break
	case "POST":
		app := &models.Applicant{
			BattleID:             r.FormValue("app_battleid"),
			Battletag:            r.FormValue("app_battletag"),
			Character:            r.FormValue("app_character"),
			Email:                r.FormValue("app_email"),
			RealName:             r.FormValue("app_realname"),
			Location:             r.FormValue("app_location"),
			Age:                  r.FormValue("app_age"),
			Gender:               r.FormValue("app_gender"),
			ComputerSpecs:        r.FormValue("app_computerspecs"),
			PreviousGuilds:       r.FormValue("app_previousguilds"),
			ReasonsLeavingGuilds: r.FormValue("app_reasonsleavingguilds"),
			WhyJoinThisGuild:     r.FormValue("app_whyjointhisguild"),
			References:           r.FormValue("app_references"),
			FinalRemarks:         r.FormValue("app_finalremarks"),
		}

		err := app.Save()

		if nil != err {
			e := &errors.Error{
				Message: err.Error(),
				Package: "api.beta",
				Type:    "API",
				Method:  "Apply",
			}
			a.Error(w, e)
			return
		}

		a.JSON(w, true)
		break
	default:
		e := &errors.Error{
			Message: "default hit",
			Package: "api.beta",
			Type:    "API",
			Method:  "Apply",
		}
		a.Error(w, e)
		return
	}
}
