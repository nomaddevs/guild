package beta

import (
	"encoding/json"
	"net/http"

	"github.com/munsy/battlenet"
	"github.com/munsy/guild/errors"
)

type API struct {
	settings     *battlenet.Settings
	url          string
	key          string
	secret       string
	authRedirect string
}

func New(s *battlenet.Settings, url, key, secret, authRedirect string) *API {
	return &API{
		settings:     s,
		url:          url,
		key:          key,
		secret:       secret,
		authRedirect: authRedirect,
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

	// Login and user data
	mux.HandleFunc(EndpointLogin, a.LoginRedirect)
	mux.HandleFunc(EndpointCallback, a.LoginCallback)
	mux.HandleFunc(EndpointUser, a.User)
}

func (a *API) JSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(v)
}

func (a *API) Error(w http.ResponseWriter, err *errors.Error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err)
}
