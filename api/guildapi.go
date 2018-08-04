package api

import (
	"net/http"

	"github.com/munsy/guild/api/beta"
)

type GuildAPI struct {
	Settings *APISettings
	Beta     *beta.API
}

func New(a *APISettings) *GuildAPI {
	return &GuildAPI{
		Settings: a,
		Beta:     beta.New(a.BlizzardSettings, a.BlizzardCallbackURL, a.Key, a.Secret, a.AuthRedirect),
	}
}

func (g *GuildAPI) Load() *http.ServeMux {
	mux := http.NewServeMux()

	g.Beta.Load(mux)

	return mux
}
