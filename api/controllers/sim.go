package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/munsy/guild/models"
)

// Sim page.
func handleSim(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "cupcake")
	if err != nil {
		fmt.Printf("Invalid session %v\n", session)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if nil != err {
		fmt.Printf("Error:" + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var user models.BnetUser
	var Characters models.Characters
	var url string
	if access_token, ok := session.Values["usercode"].(string); ok {
		// Get user's Battle.net ID and Battletag.
		url = "https://us.api.battle.net/account/user?access_token=" + access_token
		err := models.Get(url, &user)
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
			"sim",
			user,
			Characters,
		}

		t := template.Must(template.ParseFiles(home+"/views/base.html", home+"/views/libraries.html", home+"/views/navbar.html", home+"/views/sim.html"))
		t.ExecuteTemplate(w, "base", data)
		break
	default:
		fmt.Fprintln(w, "Sorry, nothing here!")
	}
}
