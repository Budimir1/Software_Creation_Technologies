package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"tech-ip-sem2/gen/proto/auth"
)

type Server struct {
	auth.UnimplementedAuthServiceServer
}

func (s *Server) Verify(ctx context.Context, req *auth.VerifyRequest) (*auth.VerifyResponse, error) {
	// Проверка токена (учебный пример)
	token := req.Token
	if token == "" {
		return nil, status.Errorf(codes.Unauthenticated, "token is empty")
	}
	// Считаем валидным только токен "valid-token"
	if token != "valid-token" {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token")
	}
	// В реальном проекте subject извлекается из токена
	return &auth.VerifyResponse{
		Valid:   true,
		Subject: "user@example.com",
	}, nil
}
