package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/munsy/battlenet"
	"github.com/munsy/guild/api"
	"github.com/munsy/guild/config"
)

const configFilename = "config.toml"

var runTLS = false

var staticRoutes = [5]string{"bootstrap", "css", "html", "images", "js"}

func redirect(w http.ResponseWriter, req *http.Request) {
	// remove/add non default ports from req.Host
	var target string

	target = "https://" + req.Host + req.URL.Path

	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}

	http.Redirect(w, req, target, http.StatusTemporaryRedirect)
}

// Create room for static files serving
func register(mux *http.ServeMux, dir string) {
	ds := "./"
	s := "/"

	dsdir := ds + dir
	sdir := s + dir
	sdirs := s + dir + s

	mux.Handle(sdirs, http.StripPrefix(sdir, http.FileServer(http.Dir(dsdir))))
}

func main() {
	fmt.Println("Starting server...")

	config.Read(configFilename)

	// Create new settings
	settings := &api.APISettings{
		BlizzardCallbackURL: "https://www.munsy.net/callback",
		BlizzardSettings: &battlenet.Settings{
			Client: &http.Client{Timeout: (10 * time.Second)},
			Locale: battlenet.Locale.AmericanEnglish,
			Region: battlenet.Regions.US,
		},
		Key:    cfg.Key,
		Secret: cfg.Secret,
	}

	guild := api.New(settings)

	mux := guild.Load()

	for i := 0; i < len(staticRoutes); i++ {
		register(mux, staticRoutes[i])
	}

	// Any other request, we should render our SPA's only html file,
	// Allowing angular to do the routing on anything else other then the api
	// and the files it needs for itself to work.
	// Order here is critical. This html should contain the base tag like
	// <base href="/"> *href here should match the HandleFunc path below
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "html/index.html")
	})

	if !runTLS {
		fmt.Println("TLS configuration not set. Falling back to HTTP...")
		http.ListenAndServe(":80", mux)
	} else {
		Addr := ":443"
		CertFile := ""
		KeyFile := ""

		fmt.Println("Redirecting HTTPS traffic to " + Addr)
		go http.ListenAndServe(":80", http.HandlerFunc(redirect))

		// Start the server through TLS/SSL.
		log.Fatal(http.ListenAndServeTLS(Addr, CertFile, KeyFile, mux))
	}
}
