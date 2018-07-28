package api

import (
	"github.com/munsy/battlenet"
)

type APISettings interface {
	BlizzardCallbackURL() string
	BlizzardSettings() *battlenet.Settings
	Key() string
	Secret() string
}
