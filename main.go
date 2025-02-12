package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	
	// Allow us to read the .env file
	godotenv.Load(".env")
	
	portString := os.Getenv("PORT")
	
	// break out of the program with error message (could not find port in our .env file)
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	router := chi.NewRouter()

	srv := &http.Server {
		Handler : router,
		Addr : ":" + portString,
	}

	log.Printf("Server is starting on port %v", portString)
	err := srv.ListenAndServe() // the code should stop here and we should handle all the http requests etc
	
	if err != nil {
		log.Fatal(err)
	}


	fmt.Println("PORT:", portString)


}