package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

// configuredUpgrader preserves same-origin access and permits explicitly listed
// browser origins, such as Cockpit on a different port. Non-browser clients
// without an Origin header remain supported.
func configuredUpgrader(originList string) (websocket.Upgrader, error) {
	allowed := make(map[string]bool)
	if strings.TrimSpace(originList) != "" {
		for _, entry := range strings.Split(originList, ",") {
			origin := strings.TrimSpace(entry)
			if !validOrigin(origin) {
				return websocket.Upgrader{}, fmt.Errorf("invalid WebSocket origin %q: expected an http(s) origin without a path", origin)
			}
			allowed[origin] = true
		}
	}
	return websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		origins := r.Header.Values("Origin")
		if len(origins) == 0 {
			return true
		}
		if len(origins) != 1 || !validOrigin(origins[0]) {
			return false
		}
		origin := origins[0]
		scheme := "http"
		if r.TLS != nil {
			scheme = "https"
		}
		return origin == scheme+"://"+r.Host || allowed[origin]
	}}, nil
}

func validOrigin(origin string) bool {
	parsed, err := url.Parse(origin)
	return err == nil && (parsed.Scheme == "http" || parsed.Scheme == "https") &&
		parsed.Host != "" && parsed.User == nil && parsed.Path == "" &&
		parsed.RawQuery == "" && !parsed.ForceQuery && parsed.Fragment == "" &&
		!strings.Contains(origin, "#")
}

func serve(address, certFile, keyFile string, handler http.Handler) error {
	if (certFile == "") != (keyFile == "") {
		return fmt.Errorf("--tls-cert and --tls-key must be provided together")
	}
	server := &http.Server{
		Addr:              address,
		Handler:           handler,
		ReadHeaderTimeout: 10 * time.Second,
		TLSConfig:         &tls.Config{MinVersion: tls.VersionTLS12},
	}
	if certFile != "" {
		return server.ListenAndServeTLS(certFile, keyFile)
	}
	return server.ListenAndServe()
}
