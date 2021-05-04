package app

import (
	"golang/microservices/controllers"
	"log"
	"net/http"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)

	log.Println("Server is starting on port 8080")
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}
