package service

import (
	"context"
	"log/slog"
	"time"

	domain "github.com/mephirious/group-project/services/auth/domain"
)

type LoggingService struct {
	logger *slog.Logger
	next   domain.Service
}

func NewLoggingService(logger *slog.Logger, next domain.Service) domain.Service {
	return &LoggingService{
		logger: logger,
		next:   next,
	}
}

func (s *LoggingService) Register(ctx context.Context, input domain.RegisterInput) (response *domain.RegisterResponse, err error) {
	start := time.Now()
	defer func() {
		logger := s.logger
		if response != nil {
			logger = s.logger.With(slog.Any("user", response.User))
		} else {
			logger = s.logger.With(slog.Any("err", err))
		}
		logger.Info(
			"Register",
			"took", time.Since(start).String(),
		)
	}()
	return s.next.Register(ctx, input)
}

func (s *LoggingService) Login(ctx context.Context, input domain.LoginInput) (response *domain.LoginResponse, err error) {
	start := time.Now()
	defer func() {
		logger := s.logger
		if response != nil {
			logger = s.logger.With(slog.Any("response", response.Message))
		} else {
			logger = s.logger.With(slog.Any("err", err))
		}
		logger.Info(
			"Login",
			"took", time.Since(start).String(),
		)
	}()
	return s.next.Login(ctx, input)
}

func (s *LoggingService) Logout(ctx context.Context, input domain.LogoutInput) (response *domain.LogoutResponse, err error) {
	start := time.Now()
	defer func() {
		logger := s.logger
		if response != nil {
			logger = s.logger.With(slog.Any("response", response.Message))
		} else {
			logger = s.logger.With(slog.Any("err", err))
		}
		logger.Info(
			"Logout",
			"took", time.Since(start).String(),
		)
	}()
	return s.next.Logout(ctx, input)
}

func (s *LoggingService) RefreshUserAccessToken(ctx context.Context, input domain.RefreshInput) (response *domain.LoginResponse, err error) {
	start := time.Now()
	defer func() {
		logger := s.logger
		if response != nil {
			logger = s.logger.With(slog.Any("response", response.Message))
		} else {
			logger = s.logger.With(slog.Any("err", err))
		}
		logger.Info(
			"Login",
			"took", time.Since(start).String(),
		)
	}()
	return s.next.RefreshUserAccessToken(ctx, input)
}

func (s *LoggingService) ValidateAccessToken(ctx context.Context, input domain.LogoutInput) (response *domain.ValidateResponse, err error) {
	start := time.Now()
	defer func() {
		logger := s.logger
		if response != nil {
			logger = s.logger.With(slog.Any("claims", response))
		} else {
			logger = s.logger.With(slog.Any("err", err))
		}
		logger.Info(
			"ValidateAccessToken",
			"took", time.Since(start).String(),
		)
	}()
	return s.next.ValidateAccessToken(ctx, input)
}
