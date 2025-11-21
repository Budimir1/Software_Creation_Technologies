package app

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"example.com/prak_9/internal/http/handlers"
	"example.com/prak_9/internal/platform/config"
	"example.com/prak_9/internal/repo"
)

func Run() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("warning: .env not loaded:", err)
	}

	cfg := config.Load()
	db, err := repo.Open(cfg.DB_DSN)
	if err != nil {
		log.Fatal("db connect:", err)
	}

	if err := db.Exec("SET timezone TO 'UTC'").Error; err != nil { /* необязательно */
	}

	users := repo.NewUserRepo(db)
	if err := users.AutoMigrate(); err != nil {
		log.Fatal("migrate:", err)
	}

	auth := &handlers.AuthHandler{Users: users, BcryptCost: cfg.BcryptCost}

	r := chi.NewRouter()
	r.Post("/auth/register", auth.Register)
	r.Post("/auth/login", auth.Login)

	log.Println("listening on", cfg.Addr)
	log.Fatal(http.ListenAndServe(cfg.Addr, r))
}
