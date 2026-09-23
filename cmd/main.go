package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "SERVER\t", log.Ldate|log.Ltime)

	srv := server.CreateRouter(logger)
	logger.Println("Сервер запускается на порту :8080...")
	err := srv.Server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}
