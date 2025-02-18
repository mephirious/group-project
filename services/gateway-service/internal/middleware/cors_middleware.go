package middleware

import (
	"log"
	"net/http"
	"os"
	"strings"
)

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		CORS_URLS := os.Getenv("CORS_URLS")
		if CORS_URLS == "" {
			log.Fatal("CORS_URLS not set in .env file")
		}

		allowedOrigins := strings.Split(CORS_URLS, ",")
		origin := r.Header.Get("Origin")

		allowOrigin := ""
		for _, o := range allowedOrigins {
			if strings.TrimSpace(o) == origin {
				allowOrigin = origin
				break
			}
		}

		if allowOrigin != "" {
			w.Header().Set("Access-Control-Allow-Origin", allowOrigin)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Cookie")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
		}

		// Handle preflight OPTIONS request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
