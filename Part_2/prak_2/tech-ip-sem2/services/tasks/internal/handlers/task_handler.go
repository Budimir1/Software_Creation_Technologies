package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"tech-ip-sem2/gen/proto/auth"
)

type TaskHandler struct {
	AuthClient auth.AuthServiceClient
}

func (h *TaskHandler) HandleTasks(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "missing token", http.StatusUnauthorized)
		return
	}

	// Устанавливаем дедлайн 2 секунды на gRPC вызов
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	log.Println("calling grpc verify")
	resp, err := h.AuthClient.Verify(ctx, &auth.VerifyRequest{Token: token})
	if err != nil {
		// Преобразуем gRPC ошибку в HTTP статус
		if st, ok := status.FromError(err); ok {
			switch st.Code() {
			case codes.Unauthenticated:
				http.Error(w, "invalid token", http.StatusUnauthorized)
			case codes.DeadlineExceeded:
				http.Error(w, "auth service timeout", http.StatusGatewayTimeout) // 504
			case codes.Unavailable:
				http.Error(w, "auth service unavailable", http.StatusServiceUnavailable) // 503
			default:
				http.Error(w, "internal error", http.StatusInternalServerError)
			}
		} else {
			http.Error(w, "unknown error", http.StatusInternalServerError)
		}
		return
	}

	if !resp.Valid {
		http.Error(w, "invalid token", http.StatusUnauthorized)
		return
	}

	// Токен валиден, формируем ответ
	response := map[string]interface{}{
		"message": "success",
		"subject": resp.Subject,
		"tasks":   []string{"task1", "task2"},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
