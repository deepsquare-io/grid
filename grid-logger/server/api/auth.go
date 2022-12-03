package api

import (
	"context"
	"errors"
	"strings"

	authv1alpha1 "github.com/deepsquare-io/the-grid/grid-logger/gen/go/auth/v1alpha1"
	"github.com/deepsquare-io/the-grid/grid-logger/logger"
	"github.com/deepsquare-io/the-grid/grid-logger/server/auth"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authAPIServer struct {
	authv1alpha1.UnimplementedAuthAPIServer
	storage     *auth.MemStorage
	jwtProvider *auth.JwtHmacProvider
}

func NewAuthAPIServer(storage *auth.MemStorage, jwtProvider *auth.JwtHmacProvider) *authAPIServer {
	if storage == nil {
		logger.I.Panic("storage is nil")
	}
	if jwtProvider == nil {
		logger.I.Panic("jwtProvider is nil")
	}
	return &authAPIServer{
		storage:     storage,
		jwtProvider: jwtProvider,
	}
}

func (s *authAPIServer) Register(ctx context.Context, req *authv1alpha1.RegisterRequest) (*authv1alpha1.RegisterResponse, error) {
	logger.I.Debug("Register", zap.Any("req", req))
	nonce, err := auth.GetNonce()
	if err != nil {
		logger.I.Error("Register.GetNonce", zap.Error(err))
		return nil, err
	}
	u := auth.User{
		Address: strings.ToLower(req.GetAddress()),
		Nonce:   nonce,
	}
	if err := s.storage.CreateIfNotExists(u); err != nil {
		logger.I.Error("Register.CreateIfNotExists", zap.Error(err))
		if errors.Is(err, auth.ErrUserExists) {
			return nil, status.Error(codes.AlreadyExists, "user already exists")
		}
		return nil, err
	}
	return &authv1alpha1.RegisterResponse{}, nil
}

func (s *authAPIServer) Nonce(ctx context.Context, req *authv1alpha1.NonceRequest) (*authv1alpha1.NonceResponse, error) {
	logger.I.Debug("Nonce", zap.Any("req", req))
	user, err := s.storage.Get(strings.ToLower(req.GetAddress()))
	if err != nil {
		logger.I.Error("Nonce.Get", zap.Error(err))
		if errors.Is(err, auth.ErrUserNotExists) {
			return nil, status.Error(codes.NotFound, "nonce not found")
		}
		return nil, err
	}
	return &authv1alpha1.NonceResponse{
		Nonce: user.Nonce,
	}, nil
}
func (s *authAPIServer) SignIn(ctx context.Context, req *authv1alpha1.SignInRequest) (*authv1alpha1.SignInResponse, error) {
	logger.I.Debug("SignIn", zap.Any("req", req))
	user, err := auth.Authenticate(s.storage, req.GetAddress(), req.GetNonce(), req.GetSig())
	if err != nil {
		logger.I.Error("SignIn.Authenticate", zap.Error(err))
		if errors.Is(err, auth.ErrAuthError) {
			return nil, status.Errorf(codes.Unauthenticated, "authentication failed with error %v", err)
		}
		return nil, err
	}
	signedToken, err := s.jwtProvider.CreateStandard(user.Address)
	if err != nil {
		logger.I.Error("SignIn.CreateStandard", zap.Error(err))
		return nil, err
	}
	return &authv1alpha1.SignInResponse{
		AccessToken: signedToken,
	}, nil
}
