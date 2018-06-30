package router

import "github.com/munsy/guild/api"

// Routes type.
type Routes []Route

// Build a function later on that parses routes.json in this directory.

// Mappings to the website, administrator panels, and other potential services.

var routes = Routes{
	// Index page routing.
	/*
		Route{"Index", "GET", "/", handleIndex},
		Route{"Index", "POST", "/", handleIndex},
		Route{"Index", "GET", "/index", handleIndex},
		Route{"Index", "POST", "/index", handleIndex},
	*/

	// Battle.net authentication routing
	//Route{"Login", "GET", "/login", handleBnetLogin},
	//Route{"Callback", "POST", "/callback", handleBnetCallback},
	//Route{"Callback", "GET", "/callback", handleBnetCallback},

	// API
	Route{"Test", "GET", api.EndpointTest, api.HandleAngular},
	Route{"Test", "POST", api.EndpointTest, api.HandleAngular},
}
