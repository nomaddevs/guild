package api

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/munsy/guild/models"
)

// Roster page
func handleRoster(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "cupcake")
	if err != nil {
		fmt.Printf("Invalid session %v\n", session)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var User models.BnetUser
	var Guildinfo models.GuildInfo
	var url string
	if access_token, ok := session.Values["usercode"].(string); ok {
		// Get user's Battle.net ID and Battletag.
		url = "https://us.api.battle.net/account/user?access_token=" + access_token
		//fmt.Printf("URL: %s\n", url)
		err := models.Get(url, &User)
		if nil != err {
			fmt.Println("Error on " + url)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	key := models.GetAPICredential(dbUsername, dbPassword, "localhost", "3306", "guild", "bnetapi", "apikey")
	url = "https://us.api.battle.net/wow/guild/thrall/NoBelfsAllowed?fields=members&locale=en_US&apikey=" + key
	err = models.Get(url, &Guildinfo)
	if nil != err {
		fmt.Println("Error on " + url)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	switch r.Method {
	case "GET":
		t, _ := template.ParseFiles(home+"/views/base.html", home+"/views/libraries.html", home+"/views/navbar.html", home+"/views/roster.html")
		t.ExecuteTemplate(w, "base", M{
			"Active":    "roster",
			"user":      User,
			"guildinfo": Guildinfo,
		})
		break
	default:
		fmt.Fprintln(w, "Sorry, nothing here!")
	}
}
