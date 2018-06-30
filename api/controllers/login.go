package controllers

import (
        "fmt"
        "net/http"

	"golang.org/x/oauth2"

)

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

