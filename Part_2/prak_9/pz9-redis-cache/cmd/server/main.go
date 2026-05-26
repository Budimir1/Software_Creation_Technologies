package main

import (
	"context"
	"log"
	"net/http"

	"budimir/pz9-redis-cache/internal/cache"
	"budimir/pz9-redis-cache/internal/config"
	"budimir/pz9-redis-cache/internal/httpapi"
	"budimir/pz9-redis-cache/internal/service"
	"budimir/pz9-redis-cache/internal/task"
)

func main() {
	cfg := config.New()

	repo := task.NewRepo()

	redisClient := cache.NewRedisClient(cfg)
	if err := cache.Ping(context.Background(), redisClient); err != nil {
		log.Println("warning: redis is unavailable at startup:", err)
	}

	taskService := service.NewTaskService(repo, redisClient, cfg)

	handler := httpapi.NewHandler(taskService)

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/tasks/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetTaskByID(w, r)
		case http.MethodPatch:
			handler.PatchTask(w, r)
		case http.MethodDelete:
			handler.DeleteTask(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("server started on :8082")
	if err := http.ListenAndServe(":8082", mux); err != nil {
		log.Fatal(err)
	}
}
