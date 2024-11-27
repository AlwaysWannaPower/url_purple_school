package main

import (
	"fmt"
	"gorut/configs"
	"gorut/internal/hell"
	"log"
	"net/http"
)

func main() {

	conf := configs.LoadConfig()
	router := http.NewServeMux()
	hell.NewHellHandler(router)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8081")
	log.Fatal(server.ListenAndServe())
}
