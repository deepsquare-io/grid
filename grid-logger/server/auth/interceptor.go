package auth

import (
	"context"
	"errors"

	"github.com/deepsquare-io/the-grid/grid-logger/logger"
	middleware "github.com/grpc-ecosystem/go-grpc-middleware/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type interceptor struct {
	jwtProvider      *JwtHmacProvider
	storage          *MemStorage
	accessibleRoutes map[string]bool
}

func NewInterceptor(
	jwtProvider *JwtHmacProvider,
	storage *MemStorage,
	accessibleRoutes map[string]bool,
) *interceptor {
	if jwtProvider == nil {
		logger.I.Panic("jwtProvider is nil")
	}
	if storage == nil {
		logger.I.Panic("storage is nil")
	}
	return &interceptor{
		jwtProvider:      jwtProvider,
		storage:          storage,
		accessibleRoutes: accessibleRoutes,
	}
}

func (i *interceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		user, err := i.authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}

		return handler(context.WithValue(ctx, UserField{}, user), req)
	}
}

func (i *interceptor) Stream() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		ctx := stream.Context()
		user, err := i.authorize(ctx, info.FullMethod)
		if err != nil {
			return err
		}
		newStream := middleware.WrapServerStream(stream)
		newStream.WrappedContext = context.WithValue(ctx, UserField{}, user)

		return handler(srv, newStream)
	}
}

func (i *interceptor) authorize(ctx context.Context, method string) (User, error) {
	ok := i.accessibleRoutes[method]
	if !ok {
		// everyone can access
		return User{}, nil
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return User{}, status.Error(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return User{}, status.Error(codes.Unauthenticated, "authorization token is not provided")
	}

	const prefix = "Bearer "
	if len(values[0]) < len(prefix) {
		return User{}, status.Errorf(codes.Unauthenticated, "access token is in bad format, expected: 'Bearer <token>', received: %s", values[0])
	}

	accessToken := values[0][len(prefix):]
	claims, err := i.jwtProvider.Verify(accessToken)
	if err != nil {
		return User{}, status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	for _, ok := range i.accessibleRoutes {
		if ok {
			user, err := i.storage.Get(claims.Subject)
			if err != nil {
				if errors.Is(err, ErrUserNotExists) {
					return User{}, status.Errorf(codes.Unauthenticated, "user not found: %v", err)
				}
				return User{}, err
			}
			return user, nil
		}
	}

	return User{}, status.Error(codes.PermissionDenied, "no permission to access this RPC")
}
