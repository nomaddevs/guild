package beta

import (
	"fmt"
	"net/http"

	"github.com/munsy/battlenet"
	//"github.com/munsy/guild/config"
	//"github.com/munsy/guild/database"
	"github.com/munsy/guild/pkg/models"
)

func (a *API) Apply(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Get user token
		// If no token, show blizz auth
		c, err := r.Cookie("token")

		if nil != err {
			a.JSON(w, err)
		}

		// Get character data
		client, err := battlenet.AccountClient(a.settings, c.Value)

		if nil != err {
			a.JSON(w, err)
		}

		// Send character data for them to choose
		response, err := client.WoWOauthProfile()

		if nil != err {
			a.JSON(w, err)
		}

		a.JSON(w, response.Data)
		break
	case "POST":
		app := &models.Applicant{
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
			a.JSON(w, err)
		}

		a.JSON(w, true)

		break
	default:
		fmt.Fprintln(w, "Sorry, nothing here!")
	}
}
