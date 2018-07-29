package beta

import (
	//"fmt"
	"net/http"

	"github.com/munsy/battlenet"
	"golang.org/x/oauth2"
)

func (a *API) LoginRedirect(w http.ResponseWriter, r *http.Request) {
	config := &oauth2.Config{
		RedirectURL:  a.url,
		ClientID:     a.key,
		ClientSecret: a.secret,
		Scopes:       []string{"wow.profile"},
		Endpoint:     battlenet.Endpoint(battlenet.Regions.US),
	}
	// Some random string, random for each request
	oauthStateString := "random"

	url := config.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

/* fix
func (a *API) LoginCallback(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "guild")

	r.ParseForm()

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
*/
