package main

import (
	"fmt"
	"log"

	"codeujuzi.github.io/database"
)

func main() {
	fmt.Println("starting server at http://localhost:8000")
	database.InitDB()

	routes()

	log.Fatal(server.ListenAndServe())
}
