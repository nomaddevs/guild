package beta

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/munsy/battlenet"
	"github.com/munsy/guild/errors"
	"github.com/munsy/guild/pkg/applicants"
)

func (a *API) Apply(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Get user token
		// If no token, show blizz auth
		c, err := r.Cookie("token")

		if nil != err {
			e := &errors.Error{
				Message: err.Error(),
				Package: "api.beta",
				Type:    "API",
				Method:  "Apply GET",
			}
			a.Error(w, e)
			return
		}

		client, err := battlenet.AccountClient(a.settings, c.Value)

		if nil != err {
			e := &errors.Error{
				Message: err.Error(),
				Package: "api.beta",
				Type:    "API",
				Method:  "Apply GET",
			}
			a.Error(w, e)
			return
		}

		// Send character data for them to choose from for app process
		response, err := client.WoWOauthProfile()

		if nil != err {
			e := &errors.Error{
				Message: err.Error(),
				Package: "api.beta",
				Type:    "API",
				Method:  "Apply GET",
			}
			a.Error(w, e)
			return
		}

		a.JSON(w, response.Data)
		break
	case "POST":
		body, err := ioutil.ReadAll(r.Body)

		if nil != err {
                        e := &errors.Error{
                                Message: err.Error(),
                                Package: "api.beta",
                                Type:    "API",
                                Method:  "Apply POST",
                        }
                        a.Error(w, e)
                        return
                }

		app := &applicants.Applicant{}

		err = json.Unmarshal([]byte(body), app)

		if nil != err {
                        e := &errors.Error{
                                Message: err.Error(),
                                Package: "api.beta",
                                Type:    "API",
                                Method:  "Apply POST",
                        }
                        a.Error(w, e)
                        return
                }

		err = app.Save()

		if nil != err {
			e := &errors.Error{
				Message: err.Error(),
				Package: "api.beta",
				Type:    "API",
				Method:  "Apply POST",
			}
			a.Error(w, e)
			return
		}

		a.JSON(w, true)
		break
	default:
		e := &errors.Error{
			Message: "default hit",
			Package: "api.beta",
			Type:    "API",
			Method:  "Apply DEFAULT",
		}
		a.Error(w, e)
		return
	}
}
