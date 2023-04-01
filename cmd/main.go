package main

import (
	Pet_project_ToDoApp "Pet-project-ToDoApp"
	"Pet-project-ToDoApp/pkg/handler"
	//"github.com/zhashkevych/todo-app"
	//"github.com/zhashkevych/todo-app/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(Pet_project_ToDoApp.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
