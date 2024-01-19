package main

import (
	"log"
	"net/http"
	"simple-go-web/config"
)

func main() {
	config.ConnectDB()

	log.Println("Server running on PORT 8000")
	http.ListenAndServe(":8080", nil)
}
