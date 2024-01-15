package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/akshayDhingra01/rss-aggregator-golang/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	fmt.Println("Hello World")

	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port not found in the environment")
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DB Url not found in the environment")
	}

	conn, err := sql.Open("postgres", dbUrl)

	if err != nil {
		log.Fatal("Can't connect to database", err)
	}

	apiCfg := apiConfig{ // api config which can be passed upon to the handler
		DB: database.New(conn),
	}

	fmt.Printf("Project is working on Port %s \n", port)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"LINK"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1router := chi.NewRouter()
	// v1router.HandleFunc("/healthz", handlerReadiness) // This will allow on all HTTP Request if want to do it in specific see below

	v1router.Get("/healthz", handlerReadiness)
	v1router.Get("/error", handleError)
	v1router.Get("/users", apiCfg.handlerCreateUser)

	router.Mount("/v1", v1router) // This is done for creating two different handlers : 1 for v1 and 1 for v2 if happens : Standard

	fmt.Printf("Server started working on Port %s \n", port)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port, // Address
	}

	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

	return
}
