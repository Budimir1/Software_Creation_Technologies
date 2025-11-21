package server

import (
	"log"
	"net/http"

	"example.com/prak_10/internal/http"
	"example.com/prak_10/internal/platform/config"
)

func main() {
	cfg := config.Load()
	mux := router.Build(cfg) // см. следующий шаг
	log.Println("listening on", cfg.Port)
	log.Fatal(http.ListenAndServe(cfg.Port, mux))
}
