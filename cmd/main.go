package main

import (
	"github.com/No-name16/InnoTaxi-User/InnoTaxi-User/internal/handler"
	"github.com/No-name16/InnoTaxi-User/InnoTaxi-User/internal/repository"
	"github.com/No-name16/InnoTaxi-User/InnoTaxi-User/internal/service"
	"log"
	"net/http"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	if err := http.ListenAndServe(":80", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
