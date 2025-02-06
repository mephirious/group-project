package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	domain "github.com/mephirious/group-project/services/auth/domain"
	_ "github.com/mephirious/group-project/services/auth/service"
)

type ApiServer struct {
	svc domain.Service
	srv *http.Server
}

func NewApiServer(svc domain.Service) *ApiServer {
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
	http.HandleFunc("GET "+prefix+"/logout", s.logoutHandler)
	http.HandleFunc("GET "+prefix+"/refresh", s.refreshHandler)
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
	var input domain.RegisterInput
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
		Path:     "/auth/api/v1/refresh",
		HttpOnly: true,
		Secure:   isSecure,
		SameSite: http.SameSiteStrictMode,
	})

	writeJSON(w, http.StatusOK, response.User)
}

func (s *ApiServer) loginHandler(w http.ResponseWriter, r *http.Request) {
	var input domain.LoginInput
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
		Path:     "/auth/api/v1/refresh",
		HttpOnly: true,
		Secure:   isSecure,
		SameSite: http.SameSiteStrictMode,
	})

	writeJSON(w, http.StatusOK, map[string]any{"message": response.Message})
}

func (s *ApiServer) logoutHandler(w http.ResponseWriter, r *http.Request) {
	// Get access token from cookies
	accessToken, err := r.Cookie("access_token")
	if err != nil {
		writeJSON(w, http.StatusUnauthorized, map[string]any{"error": "No access token provided"})
		return
	}
	response, err := s.svc.Logout(context.Background(), domain.LogoutInput{
		AccessToken: accessToken.Value,
	})
	if err != nil {
		writeJSON(w, http.StatusUnauthorized, map[string]any{"error": err.Error()})
		return
	}
	// Clear authentication cookies
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   os.Getenv("SERVICE_ENV") == "production",
		MaxAge:   -1,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Path:     "/auth/api/v1/refresh",
		HttpOnly: true,
		Secure:   os.Getenv("SERVICE_ENV") == "production",
		MaxAge:   -1,
	})

	writeJSON(w, http.StatusOK, map[string]any{"message": response.Message})
}

func (s *ApiServer) refreshHandler(w http.ResponseWriter, r *http.Request) {
	refreshToken, err := r.Cookie("refresh_token")
	if err != nil {
		writeJSON(w, http.StatusUnauthorized, map[string]any{"error": "Missing refresh token"})
		return
	}
	response, err := s.svc.RefreshUserAccessToken(context.Background(), domain.RefreshInput{
		RefreshToken: refreshToken.Value,
	})
	if err != nil {
		writeJSON(w, http.StatusUnauthorized, map[string]any{"error": err.Error()})
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

	if response.RefreshToken != "" {
		http.SetCookie(w, &http.Cookie{
			Name:     "refresh_token",
			Value:    response.RefreshToken,
			Path:     "/auth/api/v1/refresh",
			HttpOnly: true,
			Secure:   isSecure,
			SameSite: http.SameSiteStrictMode,
		})
	}

	writeJSON(w, http.StatusOK, map[string]any{"message": response.Message})
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
