package beta

import (
	"fmt"
	"net/http"

	"github.com/munsy/battlenet"
)

// RealmStatus handles realm status data
func (a *API) RealmStatus(w http.ResponseWriter, r *http.Request) {
	client, err := battlenet.WoWClient(a.settings, a.key)

	if nil != err {
		fmt.Fprintln(w, "ERROR GETTING WOWCLIENT:\n"+err.Error())
		fmt.Println(w, "ERROR GETTING WOWCLIENT:\n"+err.Error())
		return
	}

	response, err := client.RealmStatus()

	if nil != err {
		fmt.Fprintln(w, "ERROR GETTING REALM STATUS:\n"+err.Error())
		fmt.Println(w, err.Error())
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
