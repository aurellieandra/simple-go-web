package main

import (
	"log"
	"net/http"
	"simple-go-web/config"
	homecontroller "simple-go-web/controllers/HomeController"
	"simple-go-web/controllers/categorycontroller"
)

func main() {
	config.ConnectDB()

	http.HandleFunc("/", homecontroller.Welcome)

	// categories
	http.HandleFunc("/categories/", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	log.Println("Server running on PORT 8000")
	http.ListenAndServe(":8080", nil)
}
