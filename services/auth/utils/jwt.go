package utils

import (
	"encoding/json"
	"errors"
	"log/slog"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var (
	JWTSecret          = ""
	JWTRefreshSecret   = ""
	DefaultAudience    = "user"
	AccessTokenExpiry  = 15 * time.Minute
	RefreshTokenExpiry = 25 * time.Hour
)

type AccessTokenPayload struct {
	UserID    string `json:"userId"`
	SessionID string `json:"sessionId"`
	jwt.RegisteredClaims
}

type RefreshTokenPayload struct {
	SessionID string `json:"sessionId"`
	jwt.RegisteredClaims
}

type SignOptions struct {
	ExpiresIn time.Duration
	Secret    string
	Audience  string
}

var defaultOptions = SignOptions{
	ExpiresIn: AccessTokenExpiry,
	Secret:    JWTSecret,
	Audience:  DefaultAudience,
}

func init() {
	if err := godotenv.Load(); err != nil {
		slog.Error("Error loading .env file")
	}
	JWTSecret = os.Getenv("JWT_SECRET")
	JWTRefreshSecret = os.Getenv("JWT_REFRESH_SECRET")
}

func SignToken(payload interface{}, options *SignOptions) (string, error) {
	if options == nil {
		options = &defaultOptions
	}

	claims := jwt.MapClaims{
		"aud": options.Audience,
		"exp": time.Now().Add(options.ExpiresIn).Unix(),
	}

	if p, ok := payload.(map[string]interface{}); ok {
		for key, value := range p {
			claims[key] = value
		}
	} else {
		return "", errors.New("invalid payload format")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(options.Secret))
}

// VerifyToken verifies a JWT token and extracts its payload
func VerifyToken[T any](tokenString string, secret string) (*T, error) {
	if secret == "" {
		secret = JWTSecret
	}

	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {
		var payload T
		claimsBytes, err := json.Marshal(claims)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(claimsBytes, &payload)
		if err != nil {
			return nil, err
		}
		return &payload, nil
	}

	return nil, errors.New("invalid token")
}
