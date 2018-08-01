package beta

import (
	"net/http"
	"time"

	//"github.com/munsy/battlenet"
	"github.com/munsy/guild/config"
	"golang.org/x/oauth2"
)

func (a *API) LoginRedirect(w http.ResponseWriter, r *http.Request) {
	println("LoginRedirect() - Add " + r.URL.RequestURI() + " to cookie 'redirectURL'")
	expiration := time.Now().Add(1 * time.Hour)
	cookie := http.Cookie{Name: "redirectURL", Value: r.URL.RequestURI(), Expires: expiration}
	http.SetCookie(w, &cookie)

	http.Redirect(w, r, config.Oauth2.AuthCodeURL("state"), http.StatusTemporaryRedirect)
}

func (a *API) LoginCallback(w http.ResponseWriter, r *http.Request) {
	println(r.Method)

	r.ParseForm()

	token, err := config.Oauth2.Exchange(oauth2.NoContext, r.FormValue("code"))

	if nil != err {
		a.JSON(w, err)
		break
	}

	println("LoginCallback() - Add %v to cookie 'token'", token.AccessToken)
	expiration := time.Now().Add(1 * time.Hour)
	cookie := http.Cookie{Name: "token", Value: token.AccessToken, Expires: expiration}
	http.SetCookie(w, &cookie)

	c, err := r.Cookie("redirectURL")

	if nil != err {
		a.JSON(w, err)
		break
	}

	println("LoginCallback() - Redirect to " + c.Value)
	http.Redirect(w, r, c.Value, http.StatusTemporaryRedirect)
}
