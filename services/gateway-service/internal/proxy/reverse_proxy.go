package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// ReverseProxyHandler forwards requests to the target service
func ReverseProxyHandler(target string) http.HandlerFunc {
	targetURL, err := url.Parse(target)
	if err != nil {
		log.Fatalf("Invalid target URL: %v", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	return func(w http.ResponseWriter, r *http.Request) {
		r.Host = targetURL.Host
		r.URL.Scheme = targetURL.Scheme
		r.URL.Host = targetURL.Host

		// Forward request to the target service
		proxy.ServeHTTP(w, r)
	}
}
