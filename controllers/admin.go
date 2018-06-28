package controllers

import (
	"fmt"
	//        "html/template"
	"net/http"

	"github.com/munsy/guild/models"
)

func handleAdmin(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "cupcake")
	check(w, err)

	var User models.BnetUser
	var Characters models.Characters
	var Guildinfo models.GuildInfo
	var url string
	if access_token, ok := session.Values["usercode"].(string); ok {
		url = "https://us.api.battle.net/account/user?access_token=" + access_token
		err := models.Get(url, &User)
		check(w, err)

		url = "https://us.api.battle.net/wow/user/characters?access_token=" + access_token
		err = models.Get(url, &Characters)
		check(w, err)
	}

	gr, err := models.NewGuildReader()
	check(w, err)
	key, err := gr.GetBNetAPIKey()
	check(w, err)

	url = "https://us.api.battle.net/wow/guild/thrall/NoBelfsAllowed?fields=members&locale=en_US&apikey=" + key
	err = models.Get(url, &Guildinfo)
	check(w, err)

	data := struct {
		Active     string
		User       models.BnetUser
		Characters models.Characters
		CanPost    bool
	}{
		"",
		models.BnetUser{},
		models.Characters{},
		false,
	}

	switch r.Method {
	case "GET":
		if IsAdmin(Guildinfo, Characters) {
			data.Active = "admin"
			data.User = User
			data.Characters = Characters
			data.CanPost = IsAdmin(Guildinfo, Characters)
		}

		combineTpl(w, data, "admin")
		break
	case "POST":
		if IsAdmin(Guildinfo, Characters) {
			r.ParseForm()

			data.Active = "new_news_post"
			data.User = User
			data.Characters = Characters
			data.CanPost = true
		}

		combineTpl(w, data, "new_news_post")
		break
	default:
		fmt.Fprintln(w, "Sorry, nothing here!")
	}
}
