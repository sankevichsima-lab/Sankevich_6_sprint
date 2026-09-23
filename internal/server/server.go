package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
	"github.com/go-chi/chi/v5"
)

type Server struct{
	logger *log.Logger
	server *http.Server
}

func CreateRouter(l *log.Logger) (s Server){
	r := chi.NewRouter()

	r.Get("/", handlers.HandlRoot)	
	r.Post("/upload", handlers.HandleUploat)

	httpServ := &http.Server{
		Addr: "8080",
		Handler: r,
		ErrorLog: l,
		ReadTimeout: 5*time.Second,
		WriteTimeout: 10*time.Second,
		IdleTimeout: 15*time.Second,
	}

	return Server{
		logger: l,
		server: httpServ,
	}
}