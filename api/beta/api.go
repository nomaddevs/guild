package beta

import (
	"errors"
	"net/http"
)

type API struct {
	settings *APISettings
}

func New(s *APISettings) {
	return &API{
		settings: s,
	}
}

func (a *API) Load(mux *http.ServeMux) { //(*http.ServeMux, error) {
	mux.HandleFunc(EndpointRealms, a.RealmStatus)
	mux.HandleFunc(EndpointRoster, a.Roster)

	// return mux, nil
}
