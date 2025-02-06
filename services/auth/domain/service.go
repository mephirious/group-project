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
		AccessToken  string `json:"accessToken"`
		RefreshToken string `json:"refreshToken"`
	}
)

type List[T any] struct {
	Total    int64
	Elements []T
}

type Service interface {
	Register(context.Context, RegisterInput) (*RegisterResponse, error)
	Login(context.Context, LoginInput) (*LoginResponse, error)
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
