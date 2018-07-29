package beta

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/munsy/battlenet"
)

// Roster page
func (a *API) Roster(w http.ResponseWriter, r *http.Request) {
	wow, err := battlenet.WoWClient(a.settings.BlizzardSettings(), a.key)

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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response.Data)
		break
	default:
		fmt.Fprintln(w, "Sorry, nothing here!")
	}
}
