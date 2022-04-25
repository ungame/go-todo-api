package main

import (
	"flag"
	"go-todo-api/tests/todos"
	"log"
)

var baseURL string

func init() {
	flag.StringVar(&baseURL, "base_url", "http://localhost:8080", "")
	flag.Parse()
}

func main() {
	todos.Run(baseURL)
	log.Println("Integration Tests OK!")
}
