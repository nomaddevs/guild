package api

import (
	"net/http"

	"github.com/munsy/battlenet"
	"github.com/munsy/guild/api/beta"
)

type GuildAPI struct {
	Settings *battlenet.Settings
	Beta     *beta.API
}

func New(a *APISettings) *GuildAPI {
	return &GuildAPI{
		Settings: a,
		Beta:     beta.New(a.BlizzardSettings, a.BlizzardCallbackURL, a.Key, a.Secret),
	}
}

func (g *GuildAPI) Load() *http.ServeMux {
	mux := http.NewServeMux()

	Beta.Load(mux)

	return mux
}
