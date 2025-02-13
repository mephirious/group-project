package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

// ReverseProxyHandler forwards requests to the target service
func ReverseProxyHandler(target string) http.HandlerFunc {
	targetURL, err := url.Parse(target)
	if err != nil {
		log.Fatalf("Invalid target URL: %v", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	director := proxy.Director
	proxy.Director = func(r *http.Request) {
		director(r)

		// Strip "/products" prefix
		r.URL.Path = strings.TrimPrefix(r.URL.Path, "/products")

		// Ensure the proxy forwards the correct host
		r.Host = targetURL.Host
		r.URL.Host = targetURL.Host
		r.URL.Scheme = targetURL.Scheme
	}

	return func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	}
}
