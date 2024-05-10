package main

import (
	"fmt"
	"log"
	"module/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.SubscriptionHandler)
	fmt.Println("Servidor web iniciado. Acesse http://localhost:4040/")

	if err := http.ListenAndServe(":4040", nil); err != nil {
		log.Fatal(err)
	}
}
