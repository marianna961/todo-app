package main

import(
	"log"
	"github.com/marianna961/todo-app/pkg/handlers"
	"github.com/marianna961/todo-app"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error  occured while running http server %s", err.Error())
	}

}