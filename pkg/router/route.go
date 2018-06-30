package router

import (
	"fmt"
	"net/http"
)

// Route struct.
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

func (r Route) ToString() string {
	return fmt.Sprintf("%-15s %-15s %-15s", r.Name, r.Method, r.Pattern)
}
