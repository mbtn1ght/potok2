package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// http://k8s.golang-school.ru:8090/mnepryakhin/my-app

func hello(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("Hello from k8s server!!"))
	if err != nil {
		fmt.Println("Error writing response:", err)
	}

	fmt.Println("200 OK! Hello handler called")
}

func main() {
	router := chi.NewRouter()
	router.Get("/ikaoden/my-app", hello)

	fmt.Println("Starting server on :8080")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
