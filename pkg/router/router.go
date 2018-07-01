package router

import (
	"fmt"

	"github.com/gorilla/mux"
)

// NewRouter makes a new router for the API.
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		fmt.Printf("[ROUTE] Loaded %s\n", route.ToString())
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)
	}
	return router
}

func Load(router *mux.Router, r Routes) {
	for _, route := range r {
		fmt.Printf("[ROUTE] Loaded %s\n", route.ToString())
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)
	}
}
