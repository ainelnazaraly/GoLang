package main

import (
	"fmt"
	"log"
	"go-rest-api-example/pkg/chars"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/health" , characters.HealthCheck).Methods("GET")
	router.HandleFunc("/characters", characters.GetCharacters).Methods("GET")
	router.HandleFunc("/characters/{name}", characters.GetCharacterByName).Methods("GET")
	
	port := 8080
	fmt.Printf("Server is running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), router); err != nil {
		log.Fatal(err)
	}
}