package server

import (
	"chat-grpc/internal/storage"
	"context"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type AuthInterceptor struct {
	session *storage.RedisSessionStorage
}

func NewAuthInterceptor(session *storage.RedisSessionStorage) *AuthInterceptor {
	return &AuthInterceptor{session: session}
}

func (i *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		if isAuthMethod(info.FullMethod) {
			return handler(ctx, req)
		}

		userID, newCtx, err := i.authenticate(ctx)
		if err != nil {
			return nil, err
		}

		// Передаем userID через контекст
		_ = userID

		return handler(newCtx, req)
	}
}

// Stream для ChatStream
func (i *AuthInterceptor) Stream() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		ss grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		if isAuthMethod(info.FullMethod) {
			return handler(srv, ss)
		}

		userID, newCtx, err := i.authenticate(ss.Context())
		if err != nil {
			return err
		}

		// Оборачиваем Stream с новым контекстом
		wrapped := &wrappedStream{ServerStream: ss, ctx: newCtx}

		_ = userID

		return handler(srv, wrapped)
	}
}

// Проверяем токен
func (i *AuthInterceptor) authenticate(ctx context.Context) (string, context.Context, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", ctx, status.Error(codes.Unauthenticated, "missing metadata")
	}

	vals := md.Get("authorization")
	if len(vals) == 0 {
		return "", ctx, status.Error(codes.Unauthenticated, "missing authorization header")
	}

	header := vals[0]

	if !strings.HasPrefix(header, "Bearer ") {
		return "", ctx, status.Error(codes.Unauthenticated, "invalid authorization format")
	}

	token := strings.TrimPrefix(header, "Bearer ")

	// Проверка токена в Redis
	userID, err := i.session.GetUserByToken(ctx, token)
	if err != nil {
		return "", ctx, status.Error(codes.Unauthenticated, "invalid or expired token")
	}

	newCtx := context.WithValue(ctx, ctxKeyUserID, userID)

	return userID, newCtx, nil
}

func isAuthMethod(method string) bool {
	switch method {
	case "/chat.AuthService/Login",
		"/chat.AuthService/Register":
		return true
	}
	return false
}

// context key
type ctxKey string

const ctxKeyUserID ctxKey = "UserID"

// Обертка Stream с новым контекстом
type wrappedStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (w *wrappedStream) Context() context.Context {
	return w.ctx
}
