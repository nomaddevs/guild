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

func (a *API) LoginRedirect(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:    "redirectURL",
		Value:   r.URL.RequestURI(),
		Expires: time.Now().Add(1 * time.Hour),
	}

	Oauth2.ClientID = a.key
	Oauth2.ClientSecret = a.secret
	Oauth2.RedirectURL = a.authRedirect

	println("ClientID: " + Oauth2.ClientID)
	println("ClientSecret: " + Oauth2.ClientSecret)
	println("RedirectURL: " + Oauth2.RedirectURL)

	http.SetCookie(w, &cookie)
	http.Redirect(w, r, Oauth2.AuthCodeURL("state"), http.StatusTemporaryRedirect)
}

func (a *API) LoginCallback(w http.ResponseWriter, r *http.Request) {
	println(r.Method)

	r.ParseForm()

	token, err := Oauth2.Exchange(oauth2.NoContext, r.FormValue("code"))

	if nil != err {
		e := &errors.Error{
			Message: err.Error(),
			Package: "api.beta",
			Type:    "API",
			Method:  "LoginCallback Line 48",
		}
		a.Error(w, e)
		return
	}

	expiration := time.Now().Add(1 * time.Hour)
	cookie := http.Cookie{Name: "token", Value: token.AccessToken, Expires: expiration}
	http.SetCookie(w, &cookie)

	c, err := r.Cookie("redirectURL")

	if nil != err {
		e := &errors.Error{
			Message: err.Error(),
			Package: "api.beta",
			Type:    "API",
			Method:  "LoginCallback Line 65",
		}
		a.Error(w, e)
		return
	}

	http.Redirect(w, r, c.Value, http.StatusTemporaryRedirect)
}
