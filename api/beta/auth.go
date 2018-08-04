package beta

import (
	"net/http"
	"time"

	"github.com/munsy/battlenet"
	"github.com/munsy/guild/errors"
	"golang.org/x/oauth2"
)

var Oauth2 = &oauth2.Config{
	Scopes:   []string{"wow.profile"},
	Endpoint: battlenet.Endpoint(battlenet.Regions.US),
}
var authstate = "state"

func (a *API) LoginRedirect(w http.ResponseWriter, r *http.Request) {
	Oauth2.ClientID = a.key
	Oauth2.ClientSecret = a.secret
	Oauth2.RedirectURL = a.authRedirect

	http.Redirect(w, r, Oauth2.AuthCodeURL("state"), http.StatusTemporaryRedirect)
}

func (a *API) LoginCallback(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	state := r.FormValue("state")

	if state != authstate {
		e := &errors.Error{
			Message: "invalid state",
			Package: "api.beta",
			Type:    "API",
			Method:  "LoginCallback Line 43",
		}
		a.Error(w, e)
		return
	}

	code := r.FormValue("code")

	token, err := Oauth2.Exchange(oauth2.NoContext, code)

	if nil != err {
		e := &errors.Error{
			Message: err.Error(),
			Package: "api.beta",
			Type:    "API",
			Method:  "LoginCallback Line 60",
		}
		a.Error(w, e)
		return
	}

	expiration := time.Now().Add(1 * time.Hour)
	cookie := http.Cookie{Name: "token", Value: token.AccessToken, Expires: expiration}
	http.SetCookie(w, &cookie)

	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}
