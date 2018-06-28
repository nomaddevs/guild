package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/munsy/guild/models"
)

func handleApply(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "cupcake")
	if err != nil {
		fmt.Printf("Invalid session %v\n", session)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var User models.BnetUser
	var Characters models.Characters
	var url string
	if access_token, ok := session.Values["usercode"].(string); ok {
		url = "https://us.api.battle.net/account/user?access_token=" + access_token
		err := models.Get(url, &User)
		if nil != err {
			fmt.Println("Error on " + url)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		url = "https://us.api.battle.net/wow/user/characters?access_token=" + access_token
		err = models.Get(url, &Characters)
		if nil != err {
			fmt.Println("Error on " + url)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	switch r.Method {
	case "GET":
		data := struct {
			Active     string
			User       models.BnetUser
			Characters models.Characters
		}{
			"apply",
			User,
			Characters,
		}

		t := template.Must(template.ParseFiles(home+"/views/base.html", home+"/views/libraries.html", home+"/views/navbar.html", home+"/views/apply/step1.html"))
		t.ExecuteTemplate(w, "base", data)
		break
	case "POST":
		data := struct {
			Active     string
			User       models.BnetUser
			Characters models.Characters
		}{
			"apply",
			User,
			Characters,
		}

		gw, err := models.NewGuildWriter()
		if nil != err {
			fmt.Fprintln(w, "shit:\n"+err.Error())
			return
		}

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

		err = gw.CreateApplicant(app)
		if nil != err {
			fmt.Fprintln(w, "shit:\n"+err.Error())
			return
		}

		t := template.Must(template.ParseFiles(home+"/views/base.html", home+"/views/libraries.html", home+"/views/navbar.html", home+"/views/apply/step2.html"))
		t.ExecuteTemplate(w, "base", data)
		break
	default:
		fmt.Fprintln(w, "Sorry, nothing here!")
	}
}
