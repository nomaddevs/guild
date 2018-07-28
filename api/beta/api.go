package beta

import (
	"errors"
	"net/http"
)

type API struct {
	settings *APISettings
	url      string
	key      string
	secret   string
}

func New(s *battlenet.Settings, url, key, secret string) {
	return &API{
		settings: s,
		url:      url,
		key:      key,
		secret:   secret,
	}
}

func (a *API) Load(mux *http.ServeMux) { //(*http.ServeMux, error) {
	mux.HandleFunc(EndpointRealms, a.RealmStatus)
	mux.HandleFunc(EndpointRoster, a.Roster)

	// return mux, nil
}
