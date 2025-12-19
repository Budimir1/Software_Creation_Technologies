package main

import (
	"log"
	"net/http"

	httpx "github.com/Budimir/prak_12/internal/http"
	"github.com/Budimir/prak_12/internal/repo"

	_ "github.com/Budimir/prak_12/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	mem := repo.NewNoteRepoMem()

	r := httpx.NewRouter(mem)

	r.Get("/docs/*", httpSwagger.WrapHandler)

	log.Println("Server started at :8085")
	log.Fatal(http.ListenAndServe(":8085", r))
}

// Package main Notes API server.
//
// @title           Notes API
// @version         1.0
// @description     Учебный REST API для заметок (CRUD).
// @contact.name    Backend Course
// @contact.email   example@university.ru
// @BasePath        /api/v1
