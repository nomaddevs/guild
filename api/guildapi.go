package api

import (
	"net/http"

	"github.com/munsy/guild/api/beta"
)

type GuildAPI struct {
	Settings APISettings
	Beta     *beta.API
}

func New(s APISettings) *GuildAPI {
	return &GuildAPI{
		Settings: s,
		Beta:     beta.New(s),
	}
}

func Load() *http.ServeMux {
	mux := http.NewServeMux()

	Beta.Load(mux)

	return mux
}
