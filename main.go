package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tokha04/go-postgres-gorilla/router"
)

func main() {
	r := router.Router()
	fmt.Println("Starting server on the port 8000....")

	log.Fatal(http.ListenAndServe(":8000", r))
}
