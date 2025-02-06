package service

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/mephirious/group-project/services/auth/db/mongo/repository"
	"github.com/mephirious/group-project/services/auth/domain"
	"github.com/mephirious/group-project/services/auth/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthService struct {
	DB *repository.DB
}

func NewAuthService(DB *repository.DB) domain.Service {
	return &AuthService{
		DB: DB,
	}
}

func (s *AuthService) Register(ctx context.Context, input domain.RegisterInput) (*domain.RegisterResponse, error) {
	// Step 1: Validate input
	err := input.Validate()
	if err != nil {
		return nil, err
	}

	// Step 2: Check if email already exists
	existingUser, err := s.DB.GetCustomersOne(ctx, repository.GetCustomersInput{
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
	newCustomer, err := s.DB.CreateCustomer(ctx, repository.CreateCustomerInput{
		Email:    input.Email,
		Password: hashedPassword,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	// Step 5: Create a verification code
	verificationCode, err := s.DB.CreateVerificationCode(ctx, repository.CreateVerificationCodeInput{
		UserID:    newCustomer.ID,
		Type:      repository.EmailVerification,
		ExpiresAt: time.Now().AddDate(1, 0, 0),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create verification code: %v", err)
	}

	verificationURL := fmt.Sprintf("%s/email/verify/%s", os.Getenv("APP_ORIGIN"), verificationCode.ID)
	fmt.Println("Verification URL:", verificationURL)

	// Step 6: Create a session
	session, err := s.DB.CreateSession(ctx, repository.CreateSessionInput{
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
	return &domain.RegisterResponse{
		User: domain.CustomerView{
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

func (s *AuthService) Login(ctx context.Context, input domain.LoginInput) (*domain.LoginResponse, error) {
	// Step 1: Validate input
	err := input.Validate()
	if err != nil {
		return nil, err
	}

	existingUser, err := s.DB.GetCustomersOne(ctx, repository.GetCustomersInput{
		Email: &input.Email,
	})
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, fmt.Errorf("failed to check for existing email: %v", err)
	}
	if existingUser == nil {
		return nil, fmt.Errorf("user not found")
	}

	if err := utils.ComparePassword(existingUser.Password, input.Password); err != nil {
		return nil, err
	}

	session, err := s.DB.CreateSession(ctx, repository.CreateSessionInput{
		UserID:    existingUser.ID,
		UserAgent: input.UserAgent,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %v", err)
	}

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
		"userId":    existingUser.ID,
		"sessionId": session.ID,
	}, &utils.SignOptions{
		ExpiresIn: utils.AccessTokenExpiry,
		Secret:    utils.JWTSecret,
		Audience:  utils.DefaultAudience,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create access token: %v", err)
	}

	return &domain.LoginResponse{
		Message:      "Login successful",
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *AuthService) Logout(ctx context.Context, input domain.LogoutInput) (*domain.LogoutResponse, error) {
	// Verify the token
	claims, err := utils.VerifyToken[utils.AccessTokenPayload](input.AccessToken, utils.JWTSecret)
	if err != nil {
		return nil, err
	}
	// Extract session ID from token and remove session from DB
	err = s.DB.DeleteSession(ctx, claims.SessionID)
	if err != nil {
		fmt.Println("Failed to delete session:", err)
	}
	return &domain.LogoutResponse{
		Message: "Logout successfult",
	}, nil
}

func (s *AuthService) RefreshUserAccessToken(ctx context.Context, input domain.RefreshInput) (*domain.LoginResponse, error) {
	// Verify the token
	claims, err := utils.VerifyToken[utils.RefreshTokenPayload](input.RefreshToken, utils.JWTRefreshSecret)
	if err != nil {
		return nil, err
	}

	session, err := s.DB.GetSessionOne(ctx, repository.GetSessionsInput{
		ID: &claims.SessionID,
	})
	if err != nil {
		return nil, err
	}

	now := time.Now()
	if now.After(session.ExpiresAt) {
		return nil, fmt.Errorf("Session expired")
	}

	// Refresh session if it expires in the next 24 hours
	var newRefreshToken string
	if session.ExpiresAt.Sub(now) <= 1*time.Hour {
		session.ExpiresAt = now.Add(utils.RefreshTokenExpiry)
		err := s.DB.UpdateSessionExpiry(ctx, session.ID, session.ExpiresAt)
		if err != nil {
			return nil, fmt.Errorf("failed to update session expiry: %v", err)
		}

		refreshToken, err := utils.SignToken(map[string]interface{}{
			"sessionId": session.ID,
		}, &utils.SignOptions{
			ExpiresIn: utils.RefreshTokenExpiry,
			Secret:    utils.JWTRefreshSecret,
			Audience:  utils.DefaultAudience,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create refresh token: %v", err)
		}
		newRefreshToken = refreshToken
	}

	accessToken, err := utils.SignToken(map[string]interface{}{
		"userId":    session.UserID,
		"sessionId": session.ID,
	}, &utils.SignOptions{
		ExpiresIn: utils.AccessTokenExpiry,
		Secret:    utils.JWTSecret,
		Audience:  utils.DefaultAudience,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create access token: %v", err)
	}

	return &domain.LoginResponse{
		Message:      "Access token refreshed",
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
	}, nil
}
