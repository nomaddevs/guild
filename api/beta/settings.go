package beta

import (
	"github.com/munsy/battlenet"
)

type APISettings struct {
	blizzardCallbackURL string
	blizzardSettings    *battlenet.Settings
	key                 string
	secret              string
}

func (s *APISettings) BlizzardCallbackURL() string {
	return s.blizzardCallbackURL
}

func (s *APISettings) BlizzardSettings() *battlenet.Settings {
	return s.blizzardSettings
}

func (s *APISettings) Key() string {
	return s.key
}

func (s *APISettings) Secret() string {
	return s.secret
}
