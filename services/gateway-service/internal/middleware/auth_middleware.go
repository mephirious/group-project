package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
)

// UserClaims stores verified user data
type UserClaims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
}

// ValidateToken sends the access token to auth-service for validation
func ValidateToken(token string) (*UserClaims, error) {
	authServiceURL := os.Getenv("AUTH_SERVICE_URL") + "/auth/api/v1/validate-token"

	client := &http.Client{Timeout: 5 * time.Second}
	req, err := http.NewRequest("GET", authServiceURL, nil)
	if err != nil {
		return nil, err
	}
	// set request access cookie
	req.Header.Set("Cookie", fmt.Sprintf("access_token=%s", token))

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("invalid token")
	}

	var claims UserClaims
	err = json.NewDecoder(resp.Body).Decode(&claims)
	if err != nil {
		return nil, err
	}

	return &claims, nil
}

func AuthMiddleware(next http.Handler, requiredRoles map[string]string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requiredRole, exists := requiredRoles[r.Method]
		if !exists || requiredRole == "" {
			next.ServeHTTP(w, r)
			return
		}

		// Extract token from cookies
		cookie, err := r.Cookie("access_token")
		if err != nil {
			http.Error(w, "Unauthorized: missing token", http.StatusUnauthorized)
			return
		}

		// Validate token via auth-service
		claims, err := ValidateToken(cookie.Value)
		if err != nil {
			http.Error(w, "Unauthorized: invalid token", http.StatusUnauthorized)
			return
		}

		// Check if the user's role meets the required role
		if requiredRole != "" && claims.Role != requiredRole {
			http.Error(w, "Forbidden: insufficient permissions", http.StatusForbidden)
			return
		}

		// Store user info in context
		ctx := context.WithValue(r.Context(), "user", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
