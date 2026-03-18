package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"os"
	"tech-ip-sem2/gen/proto/auth"
	"tech-ip-sem2/services/tasks/internal/handlers"
)

func main() {
	authAddr := os.Getenv("AUTH_GRPC_ADDR")
	if authAddr == "" {
		authAddr = "localhost:50051"
	}

	// Создаём соединение с Auth (без шифрования для простоты)
	conn, err := grpc.Dial(authAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to auth service: %v", err)
	}
	defer conn.Close()

	authClient := auth.NewAuthServiceClient(conn)

	handler := &handlers.TaskHandler{
		AuthClient: authClient,
	}

	http.HandleFunc("/tasks", handler.HandleTasks)

	port := os.Getenv("TASKS_PORT")
	if port == "" {
		port = "8082"
	}
	log.Printf("Tasks HTTP server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
