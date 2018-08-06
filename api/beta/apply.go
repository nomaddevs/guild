package beta

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"fmt"
	"reflect"

	"github.com/munsy/battlenet"
	"github.com/munsy/guild/errors"
	"github.com/munsy/guild/pkg/models"
)

func printResult(v interface{}) {
	fields := reflect.ValueOf(v).Elem()
	fieldTypes := fields.Type()

	for i := 0; i < fields.NumField(); i++ {
		field := fields.Field(i)
		fmt.Printf("%s: %v\n", fieldTypes.Field(i).Name, field.Interface())
	}
}

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

		app := &models.Applicant{}

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

		printResult(app)
/*		app := &models.Applicant{
			Age:                  r.FormValue("Age"),
			BattleID:             appbid,
			Battletag:            r.FormValue("Battletag"),
			ComputerSpecs:        r.FormValue("ComputerSpecs"),
			Character:            r.FormValue("Character"),
			Email:                r.FormValue("Email"),
			FinalRemarks:         r.FormValue("FinalRemarks"),
			Gender:               r.FormValue("Gender"),
			Location:             r.FormValue("Location"),
			PreviousGuilds:       r.FormValue("PreviousGuilds"),
			RealName:             r.FormValue("RealName"),
			ReasonsLeavingGuilds: r.FormValue("ReasonsLeavingGuilds"),
			References:           r.FormValue("References"),
			WhyJoinThisGuild:     r.FormValue("WhyJoinThisGuild"),
		} */
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
