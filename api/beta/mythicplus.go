package beta

import (
	"net/http"
	//"github.com/munsy/battlenet"
)

// MythicPlus handles Mythic+ data
func (a *API) MythicPlus(w http.ResponseWriter, r *http.Request) {
	a.JSON(w, "Nothing here yet!")
}
