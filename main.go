package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/LucasAMachado/rss-aggregator/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_"github.com/lib/pq"
)
type apiConfig struct {
	DB *database.Queries
}

func main() {
	
	// Allow us to read the .env file
	godotenv.Load(".env")
	
	portString := os.Getenv("PORT")
	
	// break out of the program with error message (could not find port in our .env file)
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	dbURL := os.Getenv("DB_URL")

	if dbURL == "" {
		log.Fatal("DB_URL is not found in the environment")
	}

	conn, err := sql.Open("postgres", dbURL)
	
	if err != nil {
		log.Fatal("Cannot connect to database", err)
	}

	apiCfg := apiConfig {
		DB : database.New(conn),
	}

	router := chi.NewRouter()

	// setup for cors	
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins : []string{"https://*", "http://*"},
		AllowedMethods : []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders : []string{"*"},
		ExposedHeaders : []string{"LINK"},
		AllowCredentials : false,
		MaxAge : 30,
	}))

	// All of our handlers 
	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadyness)
	v1Router.Get("/err", handlerError)
	v1Router.Post("/user", apiCfg.handlerCreateUser)

	router.Mount("/v1", v1Router)

	srv := &http.Server {
		Handler : router,
		Addr : ":" + portString,
	}

	log.Printf("Server is starting on port %v", portString)
	err = srv.ListenAndServe() // the code should stop here and we should handle all the http requests etc
	
	if err != nil {
		log.Fatal(err)
	}


	fmt.Println("PORT:", portString)


}