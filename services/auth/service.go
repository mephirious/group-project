package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/mephirious/group-project/services/authorization/mongo_util"
	"github.com/mephirious/group-project/services/authorization/utils"
	"go.mongodb.org/mongo-driver/mongo"
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
		User         mongo_util.CustomerView `json:"user"`
		AccessToken  string                  `json:"accessToken"`
		RefreshToken string                  `json:"refreshToken"`
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

type Service interface {
	Register(context.Context, RegisterInput) (*RegisterResponse, error)
	Login(context.Context, LoginInput) (*LoginResponse, error)
}

type AuthService struct {
	DB *mongo_util.DB
}

func NewAuthService(DB *mongo_util.DB) Service {
	return &AuthService{
		DB: DB,
	}
}

func validate(input RegisterInput) error {
	if input.Email == "" {
		return errors.New("email is required")
	}

	emailRegex := `^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`
	if !regexp.MustCompile(emailRegex).MatchString(input.Email) {
		return errors.New("invalid email format")
	}

	if input.Password == "" {
		return errors.New("password is required")
	}

	if len(input.Password) < 6 {
		return errors.New("password must be at least 6 characters long")
	}

	return nil
}
func (s *AuthService) Register(ctx context.Context, input RegisterInput) (*RegisterResponse, error) {
	// Step 1: Validate input
	err := validate(input)
	if err != nil {
		return nil, err
	}

	// Step 2: Check if email already exists
	existingUser, err := s.DB.GetCustomersOne(ctx, mongo_util.GetCustomersInput{
		Email: &input.Email,
	})
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, fmt.Errorf("failed to check for existing email: %v", err)
	}
	if existingUser != nil {
		return nil, fmt.Errorf("email '%s' is already in use", input.Email)
	}

	// Step 3: Hash password
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %v", err)
	}

	// Step 4: Create user in the database
	newCustomer, err := s.DB.CreateCustomer(ctx, mongo_util.CreateCustomerInput{
		Email:    input.Email,
		Password: hashedPassword,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	// Step 5: Create a verification code
	verificationCode, err := s.DB.CreateVerificationCode(ctx, mongo_util.CreateVerificationCodeInput{
		UserID:    newCustomer.ID,
		Type:      mongo_util.EmailVerification,
		ExpiresAt: time.Now().AddDate(1, 0, 0),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create verification code: %v", err)
	}

	verificationURL := fmt.Sprintf("%s/email/verify/%s", os.Getenv("APP_ORIGIN"), verificationCode.ID)
	fmt.Println("Verification URL:", verificationURL)

	// Step 6: Create a session
	session, err := s.DB.CreateSession(ctx, mongo_util.CreateSessionInput{
		UserID:    newCustomer.ID,
		UserAgent: input.UserAgent,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %v", err)
	}

	// Step 7: Generate access and refresh tokens
	refreshToken, err := utils.SignToken(map[string]interface{}{
		"sessionId": session.ID,
	}, &utils.SignOptions{
		ExpiresIn: utils.AccessTokenExpiry,
		Secret:    utils.JWTRefreshSecret,
		Audience:  utils.DefaultAudience,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create refresh token: %v", err)
	}

	accessToken, err := utils.SignToken(map[string]interface{}{
		"userId":    newCustomer.ID,
		"sessionId": session.ID,
	}, &utils.SignOptions{
		ExpiresIn: utils.AccessTokenExpiry,
		Secret:    utils.JWTSecret,
		Audience:  utils.DefaultAudience,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create access token: %v", err)
	}

	// Step 8: Return the user details and tokens
	return &RegisterResponse{
		User: mongo_util.CustomerView{
			ID:       newCustomer.ID,
			Email:    newCustomer.Email,
			Verified: newCustomer.Verified,
			// FirstName: newCustomer.FirstName,
			// LastName:  newCustomer.LastName,
			// Phone:     newCustomer.Phone,
			CreatedAt: newCustomer.CreatedAt,
			UpdatedAt: newCustomer.UpdatedAt,
		},
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *AuthService) Login(ctx context.Context, input LoginInput) (*LoginResponse, error) {
	_, err := s.DB.GetCustomersMany(ctx, mongo_util.GetCustomersInput{
		Limit:  100,
		Offset: 0,
	})
	if err != nil {
		return nil, err
	}

	return &LoginResponse{}, nil
}
