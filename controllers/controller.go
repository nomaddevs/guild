package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/mitchellh/go-bnet"
	"github.com/munsylol/guild/models"
	"golang.org/x/oauth2"
)

var home, _ = os.Getwd()

var store = sessions.NewCookieStore(securecookie.GenerateRandomKey(64))

var dbUsername = "guild" // guild
var dbPassword = "a" // a

var (
	bnetOauthConfig = &oauth2.Config{
		RedirectURL:  "https://www.munsy.net/callback",
		ClientID:     GetAPICredential(dbUsername, dbPassword, "localhost", "3306", "guild", "bnetapi", "apikey"),
		ClientSecret: GetAPICredential(dbUsername, dbPassword, "localhost", "3306", "guild", "bnetapi", "apisecret"),
		Scopes:       []string{"wow.profile"},
		Endpoint:     bnet.Endpoint("us"),
	}
	// Some random string, random for each request
	oauthStateString = "random"
)

type M map[string]interface{}

// Index page.
func handleIndex(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "cupcake")
	if err != nil {
		fmt.Printf("Invalid session %v\n", session)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newsposts, err := GetNewsPosts(dbUsername, dbPassword)
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
			Active string
			User models.BnetUser
			NewsPosts []models.NewsPost
		}{
			"home",
			user,
			newsposts,
		}
		fmt.Println("Data:", data)
		t := template.Must(template.ParseFiles(home+"/views/base.html", home+"/views/libraries.html", home+"/views/navbar.html", home+"/views/news.html"))
		t.ExecuteTemplate(w, "base", data)
		break
	default:
		fmt.Fprintln(w, "Sorry, nothing here!")
	}
}


// About page.
func handleAbout(w http.ResponseWriter, r *http.Request) {
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

        switch r.Method {
        case "GET":
                data := struct {
                        Active string
                        User models.BnetUser
                }{
                        "about",
                        user,
                }
                fmt.Println("Data:", data)
                t := template.Must(template.ParseFiles(home+"/views/base.html", home+"/views/libraries.html", home+"/views/navbar.html", home+"/views/about.html"))
                t.ExecuteTemplate(w, "base", data)
                break
        default:
                fmt.Fprintln(w, "Sorry, nothing here!")
        }
}

// Media page.
func handleMedia(w http.ResponseWriter, r *http.Request) {
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

        switch r.Method {
        case "GET":
                data := struct {
                        Active string
                        User models.BnetUser
                }{
                        "media",
                        user,
                }
                fmt.Println("Data:", data)
                t := template.Must(template.ParseFiles(home+"/views/base.html", home+"/views/libraries.html", home+"/views/navbar.html", home+"/views/media.html"))
                t.ExecuteTemplate(w, "base", data)
                break
        default:
                fmt.Fprintln(w, "Sorry, nothing here!")
        }
}

// Roster page.
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

	key := GetAPICredential(dbUsername, dbPassword, "localhost", "3306", "guild", "bnetapi", "apikey")
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

// Application page.
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
			Active string
                        User models.BnetUser
			Characters models.Characters
                }{
			"apply",
                        User,
			Characters,
                }
                fmt.Println("Data:", data)
                t := template.Must(template.ParseFiles(home+"/views/base.html", home+"/views/libraries.html", home+"/views/navbar.html", home+"/views/apply/step1.html"))
                t.ExecuteTemplate(w, "base", data)
                break
	case "POST":
		data := struct {
			Active string
                        User models.BnetUser
			Characters models.Characters
                }{
			"apply",
                        User,
			Characters,
                }
                fmt.Println("Data:", data)
                t := template.Must(template.ParseFiles(home+"/views/base.html", home+"/views/libraries.html", home+"/views/navbar.html", home+"/views/apply/step2.html"))
                t.ExecuteTemplate(w, "base",  data)
		break
	default:
		fmt.Fprintln(w, "Sorry, nothing here!")
	}
}

func handleBnetLogin(w http.ResponseWriter, r *http.Request) {
	url := bnetOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleBnetCallback(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cupcake")

	r.ParseForm()
	//fmt.Println(r.Form)
	//fmt.Println("Callback path: " + r.URL.Path)
	//fmt.Println("Callback scheme: " + r.URL.Scheme)

	switch r.Method {
	case "GET":
		state := r.FormValue("state")
		if state != oauthStateString {
			fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
			http.Redirect(w, r, "/apply", http.StatusTemporaryRedirect)
			return
		}

		code := r.FormValue("code")

		token, err := bnetOauthConfig.Exchange(oauth2.NoContext, code)
		if err != nil {
			msg := fmt.Sprintf("oauthConf.Exchange() failed with '%s'", err)
			http.Error(w, msg, http.StatusInternalServerError)
			return
		}

		//fmt.Println("Access Token: " + token.AccessToken)
		session.Values["usercode"] = token.AccessToken
		session.Save(r, w)

		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		break
	case "POST":
		fmt.Fprintln(w, "POST is working, Munsy...")
		break
	default:
		fmt.Fprintln(w, "Sorry, nothing here!")
	}
}

