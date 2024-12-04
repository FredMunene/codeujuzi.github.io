package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("starting server at http://localhost:8000")

	

	routes()

	log.Fatal(server.ListenAndServe())
}
