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

var staticRoutes = [5]string{"bootstrap", "css", "html", "images", "js"}
var guild *api.GuildAPI
var runTLS bool

func redirect(w http.ResponseWriter, req *http.Request) {
	// remove/add non default ports from req.Host
	var target string

	target = "https://" + req.Host + req.URL.Path

	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}

	http.Redirect(w, req, target, http.StatusTemporaryRedirect)
}

// Load static routes
func loadStatic(mux *http.ServeMux, dir string) {
	web := "../../web/"
	s := "/"

	webdir := web + dir
	sdir := s + dir
	sdirs := s + dir + s

	mux.Handle(sdirs, http.StripPrefix(sdir, http.FileServer(http.Dir(webdir))))
}

// Register all static routes
func register(mux *http.ServeMux) {
	for i := 0; i < len(staticRoutes); i++ {
		loadStatic(mux, staticRoutes[i])
	}

	// Any other request, we should render our SPA's only html file,
	// Allowing angular to do the routing on anything else other then the api
	// and the files it needs for itself to work.
	// Order here is critical. This html should contain the base tag like
	// <base href="/"> *href here should match the HandleFunc path below
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../../web/html/index.html")
	})
}

func init() {
	config.Read()

	runTLS = config.Addr != "" && config.CertFile != "" && config.KeyFile != ""

	settings := &api.APISettings{
		BlizzardCallbackURL: config.RedirectURL,
		BlizzardSettings: &battlenet.Settings{
			Client: &http.Client{Timeout: (10 * time.Second)},
			Locale: battlenet.Locale.AmericanEnglish,
			Region: battlenet.Regions.US,
		},
		Key:          config.Key,
		Secret:       config.Secret,
		AuthRedirect: config.RedirectURL,
	}

	guild = api.New(settings)
}

func main() {
	fmt.Println("Starting server...")

	// Load API
	mux := guild.Load()

	// Register all static routes, including index
	register(mux)

	// Any other request, we should render our SPA's only html file,
	// Allowing angular to do the routing on anything else other then the api
	// and the files it needs for itself to work.
	// Order here is critical. This html should contain the base tag like
	// <base href="/"> *href here should match the HandleFunc path below
	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "../../web/html/index.html")
	// })

	if !runTLS {
		fmt.Println("TLS configuration not set. Falling back to HTTP...")
		http.ListenAndServe(":80", mux)
	} else {
		fmt.Println("Redirecting HTTPS traffic to " + config.Addr)
		go http.ListenAndServe(":80", http.HandlerFunc(redirect))

		// Start the server through TLS/SSL.
		log.Fatal(http.ListenAndServeTLS(config.Addr, config.CertFile, config.KeyFile, mux))
	}
}
