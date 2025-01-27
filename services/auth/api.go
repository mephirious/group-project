package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type ApiServer struct {
	svc Service
	srv *http.Server
}

func NewApiServer(svc Service) *ApiServer {
	return &ApiServer{
		svc: svc,
	}
}

func (s *ApiServer) Start(listenAddr string, prefix string) error {
	s.srv = &http.Server{
		Addr: listenAddr,
	}

	http.HandleFunc("POST "+prefix+"/register", s.registerHandler)
	http.HandleFunc("POST "+prefix+"/login", s.loginHandler)
	// http.HandleFunc(prefix+"/logout", s.logoutHandler)
	// http.HandleFunc(prefix+"/refresh", s.refreshHandler)
	// http.HandleFunc(prefix+"/email/verify/{verification_code}", s.verifyEmailHandler)
	return s.srv.ListenAndServe()
}

func (s *ApiServer) Stop(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := s.srv.Shutdown(ctx); err != nil {
		log.Println("HTTP server shutdown failed:", err)
	} else {
		log.Println("Server stopped gracefully")
	}
}

func (s *ApiServer) registerHandler(w http.ResponseWriter, r *http.Request) {
	var input RegisterInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": "Invalid JSON body"})
		return
	}

	input.UserAgent = r.UserAgent()

	response, err := s.svc.Register(context.Background(), input)
	if err != nil {
		writeJSON(w, http.StatusConflict, map[string]any{"error": err.Error()})
		return
	}

	isSecure := os.Getenv("SERVICE_ENV") == "production"

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    response.AccessToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   isSecure,
		SameSite: http.SameSiteStrictMode,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    response.RefreshToken,
		Path:     "/api/v1/refresh",
		HttpOnly: true,
		Secure:   isSecure,
		SameSite: http.SameSiteStrictMode,
	})

	writeJSON(w, http.StatusOK, response.User)
}

func (s *ApiServer) loginHandler(w http.ResponseWriter, r *http.Request) {
	var input LoginInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": "Invalid JSON body"})
		return
	}

	input.UserAgent = r.UserAgent()

	response, err := s.svc.Login(context.Background(), input)
	if err != nil {
		// TODO: Gracefull error handling
		writeJSON(w, http.StatusConflict, map[string]any{"error": err.Error()})
		return
	}

	isSecure := os.Getenv("SERVICE_ENV") == "production"

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    response.AccessToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   isSecure,
		SameSite: http.SameSiteStrictMode,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    response.RefreshToken,
		Path:     "/api/v1/refresh",
		HttpOnly: true,
		Secure:   isSecure,
		SameSite: http.SameSiteStrictMode,
	})

	writeJSON(w, http.StatusOK, "Login successful")
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
