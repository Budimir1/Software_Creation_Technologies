package server

import (
	"log"
	"net/http"

	router "Budimir/prak_10/internal/http"
	"Budimir/prak_10/internal/platform/config"
)

func main() {
	cfg := config.Load()

	mux := router.Build(cfg)

	log.Println("listening on", cfg.Port)
	if err := http.ListenAndServe(cfg.Port, mux); err != nil {
		log.Fatal(err)
	}
}
