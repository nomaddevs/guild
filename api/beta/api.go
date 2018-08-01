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
	mux.HandleFunc(EndpointNews, a.News)
	mux.HandleFunc(EndpointRealms, a.RealmStatus)
	mux.HandleFunc(EndpointAbout, a.About)
	mux.HandleFunc(EndpointMedia, a.Media)
	mux.HandleFunc(EndpointRoster, a.Roster)
	mux.HandleFunc(EndpointApply, a.Apply)
	mux.HandleFunc(EndpointRecruitment, a.Recruitment)
	mux.HandleFunc(EndpointProgression, a.Progression)
	mux.HandleFunc(EndpointMythicPlus, a.MythicPlus)

	// Login
	mux.HandleFunc(EndpointLogin, a.LoginRedirect)
	mux.HandleFunc(EndpointCallback, a.LoginCallback)

	// return mux, nil
}

func (a *API) JSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
