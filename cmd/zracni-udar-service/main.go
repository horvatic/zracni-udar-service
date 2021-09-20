package main

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/horvatic/zracni-udar-service/pkg/router"
)

func main() {

	router := router.GetRoutes()

	originsOk := handlers.AllowedOrigins([]string{"http://10.0.0.100:3000"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

	http.ListenAndServe(":8080", handlers.CORS(originsOk, methodsOk)(router))
}
