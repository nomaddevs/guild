package beta

import (
	"encoding/json"
	"net/http"

	"github.com/munsy/battlenet"
)

type API struct {
	settings *battlenet.Settings
	url      string
	key      string
	secret   string
}

func New(s *battlenet.Settings, url, key, secret string) *API {
	return &API{
		settings: s,
		url:      url,
		key:      key,
		secret:   secret,
	}
}

func (a *API) Load(mux *http.ServeMux) {
	mux.HandleFunc(EndpointRealms, a.RealmStatus)
	mux.HandleFunc(EndpointRoster, a.Roster)

	// return mux, nil
}

func (a *API) JSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(v)
}
