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
		fmt.Println("ERROR GETTING WOWCLIENT:\n" + err.Error())
		a.JSON(w, err)
		return
	}

	response, err := client.RealmStatus()

	if nil != err {
		fmt.Println("ERROR GETTING REALM STATUS:\n" + err.Error())
		a.JSON(w, err)
		return
	}

	switch r.Method {
	case "GET":
		a.JSON(w, response.Data)
		break
	default:
		a.JSON(w, "Sorry, nothing here!")
	}
}
