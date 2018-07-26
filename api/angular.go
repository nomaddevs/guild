package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/munsy/gobattlenet"
)

// HandleAngular handles AngularJS
func HandleAngular(w http.ResponseWriter, r *http.Request) {
	client, err := battlenet.WoWClient(nil, "kx4h4q9xgermtsahh6n5jacpver4juzd")

	if nil != err {
		fmt.Fprintln(w, err.Error())
		return
	}

	switch r.Method {
	case "GET":
		response, err := client.RealmStatus()

		if nil != err {
			fmt.Fprintln(w, err.Error())
			return
		}

		data := response.Data

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(data)
		break
	default:
		fmt.Fprintln(w, "Sorry, nothing here!")
	}
}
