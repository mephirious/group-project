package main

import (
	"context"
	"log/slog"
	"time"
)

type LoggingService struct {
	logger *slog.Logger
	next   Service
}

func NewLoggingService(logger *slog.Logger, next Service) Service {
	return &LoggingService{
		logger: logger,
		next:   next,
	}
}

func (s *LoggingService) Register(ctx context.Context, input RegisterInput) (response *RegisterResponse, err error) {
	start := time.Now()
	defer func() {
		logger := s.logger
		if response != nil {
			logger = s.logger.With(slog.Any("user", response.User))
		}
		logger.Info(
			"Register",
			"err", err,
			"took", time.Since(start).String(),
		)
	}()
	return s.next.Register(ctx, input)
}

func (s *LoggingService) Login(ctx context.Context, input LoginInput) (response *LoginResponse, err error) {
	start := time.Now()
	defer func() {
		s.logger.Info(
			"Login",
			"response", "Login successful",
			"err", err,
			"took", time.Since(start).String(),
		)
	}()
	return s.next.Login(ctx, input)
}
