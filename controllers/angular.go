package controllers

import (
	"encoding/json"
	//"fmt"
	//"html/template"
	"net/http"
	//"github.com/munsylol/guild/models"
)

// About page
func handleAngular(w http.ResponseWriter, r *http.Request) {
	/*
		session, err := store.Get(r, "cupcake")
		if err != nil {
			fmt.Printf("Invalid session %v\n", session)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var user models.BnetUser
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
		}
	*/

	data := struct {
		Name  string
		Class string
		Spec  string
	}{
		"Munsy",
		"Druid",
		"Resto",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)

	/*
		switch r.Method {
		case "GET":
			data := struct {
				Active string
				User   models.BnetUser
			}{
				"about",
				user,
			}

			t := template.Must(template.ParseFiles(home+"/views/base.html", home+"/views/libraries.html", home+"/views/navbar.html", home+"/views/about.html"))
			t.ExecuteTemplate(w, "base", data)
			break
		default:
			fmt.Fprintln(w, "Sorry, nothing here!")
		}
	*/
}
