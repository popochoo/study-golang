package main

import (
	"concurrency/feature2"
	"concurrency/http"
	"concurrency/todo"
	"fmt"
)

func main() {
	todoList := todo.NewList()
	httpHandlers := http.NewHTTPHandlers(todoList)
	httpServer := http.NewHTTPServer(httpHandlers)

	feature2.Feature2()

	if err := httpServer.StartServer(); err != nil {
		fmt.Println("failed to start http server:", err)
	}
}
