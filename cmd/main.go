package main

import (
	"log"

	"github.com/genius321/todo-app"
	"github.com/genius321/todo-app/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRouts()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
