package main

import (
	"github.com/No-name16/InnoTaxi-User/InnoTaxi-User/internal/handler"
	"log"
	"net/http"
)

func main() {
	handlers := new(handler.Handler)
	if err := http.ListenAndServe(":80", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
