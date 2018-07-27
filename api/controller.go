package api

import (
	"fmt"
	"html/template"
	"net/http"
	//"os"
	"strings"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/mitchellh/go-bnet"
	"github.com/munsy/guild/models"
	"golang.org/x/oauth2"
)

const CAN_MAKE_NEWS_POSTS = 3

var store = sessions.NewCookieStore(securecookie.GenerateRandomKey(64))

var dbUsername = "guild" // guild
var dbPassword = "a"     // a

var (
	bnetOauthConfig = &oauth2.Config{
		RedirectURL:  "https://www.munsy.net/callback",
		ClientID:     models.GetAPICredential(dbUsername, dbPassword, "localhost", "3306", "guild", "bnetapi", "apikey"),
		ClientSecret: models.GetAPICredential(dbUsername, dbPassword, "localhost", "3306", "guild", "bnetapi", "apisecret"),
		Scopes:       []string{"wow.profile"},
		Endpoint:     bnet.Endpoint("us"),
	}
	// Some random string, random for each request
	oauthStateString = "random"
)

type M map[string]interface{}

func combineTpl(w http.ResponseWriter, data interface{}, tplName string) {
	if !strings.Contains(tplName, ".") {
		tplName += ".html"
	}
	t := template.Must(template.ParseFiles(home+"/views/base.html",
		home+"/views/libraries.html",
		home+"/views/navbar.html",
		home+"/views/"+tplName))
	t.ExecuteTemplate(w, "base", data)
}

func IsAdmin(Guildinfo models.GuildInfo, Characters models.Characters) bool {
	highestRank := -1
	for _, element := range Guildinfo.Members {
		for _, e := range Characters.CharacterList {
			if element.Character.Name == e.Name && element.Character.Realm == e.Realm && element.Character.Battlegroup == e.Battlegroup {
				fmt.Println(e.Name, e.Realm, e.Battlegroup)

				if -1 != highestRank || element.Rank < highestRank {
					highestRank = element.Rank

				}
			}
		}
	}
	return highestRank > CAN_MAKE_NEWS_POSTS
}

func check(w http.ResponseWriter, err error) {
	if nil != err {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
