package main

import (
	//"crypto/tls"
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	//"runtime"
	"encoding/json"
	"strings"
	"syscall"
	"time"

	"github.com/munsy/guild/api"
	"github.com/munsy/guild/config"
	"golang.org/x/crypto/ssh/terminal"
)

func redirect(w http.ResponseWriter, req *http.Request) {
	// remove/add non default ports from req.Host
	var target string

	target = "https://" + req.Host + req.URL.Path

	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}

	http.Redirect(w, req, target,
		http.StatusTemporaryRedirect)
}

func main() {
	fmt.Println("Starting server...")

	// Create new settings
	settings := &api.APISettings{
		BlizzardCallbackURL: "https://www.munsy.net/callback",
		BlizzardSettings: &battlenet.Settings{
			Client: &http.Client{Timeout: (10 * time.Second)},
			Locale: battlenet.Locale.AmericanEnglish,
			Region: battlenet.Regions.US,
		},
		Key:    "placeholder",
		Secret: "placeholder",
	}

	guild := api.New(settings)

	mux := guild.Load()

	// Create room for static files serving
	mux.Handle("/bootstrap/", http.StripPrefix("/bootstrap", http.FileServer(http.Dir("./bootstrap"))))
	mux.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("./css"))))
	mux.Handle("/html/", http.StripPrefix("/html", http.FileServer(http.Dir("./html"))))
	mux.Handle("/images/", http.StripPrefix("/images", http.FileServer(http.Dir("./images"))))
	mux.Handle("/js/", http.StripPrefix("/js", http.FileServer(http.Dir("./js"))))

	// Any other request, we should render our SPA's only html file,
	// Allowing angular to do the routing on anything else other then the api
	// and the files it needs for itself to work.
	// Order here is critical. This html should contain the base tag like
	// <base href="/"> *href here should match the HandleFunc path below
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "html/index.html")
	})

	if nil == cfg.TLS {
		fmt.Println("TLS configuration not set. Falling back to HTTP...")
		http.ListenAndServe(":80", mux)
	} else {
		fmt.Println("Redirecting HTTPS traffic to " + cfg.TLS.Addr)
		// Redirect all HTTP requests to HTTPS.
		go http.ListenAndServe(":80", http.HandlerFunc(redirect))

		// Start the server through TLS/SSL.
		log.Fatal(http.ListenAndServeTLS(cfg.TLS.Addr, cfg.TLS.CertFile, cfg.TLS.KeyFile, mux))
	}
}
