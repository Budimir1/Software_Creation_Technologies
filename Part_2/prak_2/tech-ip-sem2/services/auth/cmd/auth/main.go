package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	authgrpc "tech-ip-sem2/services/auth/internal/grpc"

	"google.golang.org/grpc"
	"tech-ip-sem2/gen/proto/auth"
)

func main() {
	port := os.Getenv("AUTH_GRPC_PORT")
	if port == "" {
		port = "50051"
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	authService := &authgrpc.Server{}
	auth.RegisterAuthServiceServer(s, authService)

	// Graceful shutdown
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<-sigCh
		log.Println("Shutting down gRPC server...")
		s.GracefulStop()
	}()

	log.Printf("Auth gRPC server listening on port %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
