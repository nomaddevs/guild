package beta

import (
	"fmt"
	"net/http"
)

func handleApply(w http.ResponseWriter, r *http.Request) {
	/*
		app := &models.AppInfo{
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
	*/
	fmt.Fprintln(w, "Sorry, nothing here!")
}
