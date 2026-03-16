package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	authHttp "tech-ip-sem2/services/auth/internal/http"
	"tech-ip-sem2/shared/middleware"
)

func main() {
	port := os.Getenv("auth_port")
	if port == "" {
		port = "8081"
	}

	r := mux.NewRouter()
	r.Use(middleware.RequestIDMiddleware)
	r.Use(middleware.LoggingMiddleware)

	authHttp.RegisterRoutes(r)

	addr := ":" + port
	log.Printf("auth service starting on %s", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal(err)
	}
}
