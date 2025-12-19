package main

import (
	"github.com/Budimir/prak_11/internal/http"
	"github.com/Budimir/prak_11/internal/http/handlers"
	"github.com/Budimir/prak_11/internal/repo"
	"log"
	"net/http"
)

func main() {
	repo := repo.NewNoteRepoMem()
	h := &handlers.Handler{Repo: repo}
	r := httpx.NewRouter(h)

	log.Println("Server started at :8085")
	log.Fatal(http.ListenAndServe(":8085", r))
}
