package router

import (
	"fmt"

	"github.com/gorilla/mux"
)

type Router struct {
	router *mux.Router
}

// NewRouter makes a new router for the API.
func NewRouter() *Router {
	return mux.NewRouter().StrictSlash(true)
}

/*
func Load(router *mux.Router, r Routes) {
	for _, route := range r {
		fmt.Printf("[ROUTE] Loaded %s\n", route.ToString())
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)
	}
}
*/

func (r *Router) Load(routes Routes) {
	for _, route := range routes {
		fmt.Printf("[ROUTE] Loaded %s\n", route.ToString())
		r.router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)
	}
}
