package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello World")

	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port not found in the environment")
	}

	fmt.Printf("Project is working on Port %s \n", port)

	router := chi.NewRouter()

	router.Use( cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST" , "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"LINK"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	fmt.Printf("Server started working on Port %s \n", port)

	srv := &http.Server {
		Handler : router,
		Addr : ":" + port, // Address 
	}

	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}


	return
}
