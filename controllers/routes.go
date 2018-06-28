package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
)

// Route struct.
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes type.
type Routes []Route

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

// NewRouter makes a new router for the API.
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)
	}
	return router
}

// Build a function later on that parses routes.json in this directory.

// Mappings to the website, administrator panels, and other potential services.
var routes = Routes{
	// Index page routing.
	Route{"Index", "GET", "/", handleIndex},
	Route{"Index", "POST", "/", handleIndex},
	Route{"Index", "GET", "/index", handleIndex},
	Route{"Index", "POST", "/index", handleIndex},
	Route{"Roster", "GET", "/roster", handleRoster},
	Route{"About", "GET", "/about", handleAbout},
	Route{"Media", "GET", "/media", handleMedia},

	// Battle.net authentication routing
	Route{"Login", "GET", "/login", handleBnetLogin},
	Route{"Callback", "POST", "/callback", handleBnetCallback},
	Route{"Callback", "GET", "/callback", handleBnetCallback},

	// Recruitment application routing
	Route{"Apply", "GET", "/apply", handleApply},
	Route{"Apply", "POST", "/apply", handleApply},
}
