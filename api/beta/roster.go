package beta

import (
	"fmt"
	"net/http"

	"github.com/munsy/battlenet"
)

// Roster page
func (a *API) Roster(w http.ResponseWriter, r *http.Request) {
	wow, err := battlenet.WoWClient(a.settings, a.key)

	if nil != err {
		fmt.Println(w, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := wow.RealmStatus()

	if nil != err {
		fmt.Println(w, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	switch r.Method {
	case "GET":
		a.JSON(w, response.Data)
		break
	default:
		fmt.Fprintln(w, "Sorry, nothing here!")
	}
}
