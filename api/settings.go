package api

import (
	"github.com/munsy/battlenet"
)

type APISettings struct {
	BlizzardCallbackURL string
	BlizzardSettings    *battlenet.Settings
	Key                 string
	Secret              string
	AuthRedirect        string
}
