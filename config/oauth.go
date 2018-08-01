package config

import (
	"github.com/munsy/battlenet"
	"golang.org/x/oauth2"
)

var Oauth2 = &oauth2.Config{
	ClientID:     Key,
	ClientSecret: Secret,
	RedirectURL:  RedirectURL,
	Scopes:       []string{"wow.profile"},
	Endpoint:     battlenet.Endpoint(battlenet.Regions.US),
}
