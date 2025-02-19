package domain

import (
	"context"
	"errors"
	"regexp"
)

type (
	RegisterInput struct {
		Email     string  `json:"email"`
		Password  string  `json:"password"`
		FirstName *string `json:"first_name"`
		LastName  *string `json:"last_name"`
		UserAgent string
	}
	RegisterResponse struct {
		User         CustomerView `json:"user"`
		AccessToken  string       `json:"accessToken"`
		RefreshToken string       `json:"refreshToken"`
	}
	LoginInput struct {
		Email     string `json:"email"`
		Password  string `json:"password"`
		UserAgent string
	}
	LoginResponse struct {
		Message      string `json:"message"`
		AccessToken  string `json:"accessToken"`
		RefreshToken string `json:"refreshToken"`
	}
	LogoutInput struct {
		AccessToken string
	}
	LogoutResponse struct {
		Message string `json:"message"`
	}
	RefreshInput struct {
		RefreshToken string
	}
	ValidateResponse struct {
		UserID string `json:"user_id"`
		Role   string `json:"role"`
	}
)

type List[T any] struct {
	Total    int64
	Elements []T
}

type Service interface {
	Register(context.Context, RegisterInput) (*RegisterResponse, error)
	Login(context.Context, LoginInput) (*LoginResponse, error)
	Logout(context.Context, LogoutInput) (*LogoutResponse, error)
	RefreshUserAccessToken(context.Context, RefreshInput) (*LoginResponse, error)
	ValidateAccessToken(context.Context, LogoutInput) (*ValidateResponse, error)
}

func (i *LoginInput) Validate() error {
	if i.Email == "" {
		return errors.New("email is required")
	}

	emailRegex := `^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`
	if !regexp.MustCompile(emailRegex).MatchString(i.Email) {
		return errors.New("invalid email format")
	}

	if i.Password == "" {
		return errors.New("password is required")
	}

	if len(i.Password) < 6 {
		return errors.New("password must be at least 6 characters long")
	}

	return nil
}

func (i *RegisterInput) Validate() error {
	if i.Email == "" {
		return errors.New("email is required")
	}

	emailRegex := `^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`
	if !regexp.MustCompile(emailRegex).MatchString(i.Email) {
		return errors.New("invalid email format")
	}

	if i.Password == "" {
		return errors.New("password is required")
	}

	if len(i.Password) < 6 {
		return errors.New("password must be at least 6 characters long")
	}

	return nil
}
