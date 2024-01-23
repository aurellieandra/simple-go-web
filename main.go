package main

import (
	"log"
	"net/http"
	"simple-go-web/config"
	"simple-go-web/controllers/categorycontroller"
	"simple-go-web/controllers/homecontroller"
	"simple-go-web/controllers/productcontroller"
)

func main() {
	config.ConnectDB()

	http.HandleFunc("/", homecontroller.Welcome)

	// categories
	http.HandleFunc("/categories/", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	// products
	http.HandleFunc("/products/", productcontroller.Index)
	http.HandleFunc("/products/detail", productcontroller.Detail)
	http.HandleFunc("/products/add", productcontroller.Add)
	http.HandleFunc("/products/edit", productcontroller.Edit)
	http.HandleFunc("/products/delete", productcontroller.Delete)

	log.Println("Server running on PORT 8000")
	http.ListenAndServe(":8080", nil)
}
