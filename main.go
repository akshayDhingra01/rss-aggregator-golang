package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello World")
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port not found in the environment")
	}

	fmt.Printf("Project is working on Port %s \n", port)

	return
}
