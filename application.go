package main

import (
	"RedSaludAtv/atv"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
)

func main() {
	Routes()
}

func Routes() *mux.Router {
	router := mux.NewRouter()
	atv.Routes(router)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*", "http://localhost:8000"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*", "Authorization", "Content-Type", "Access-Control-Allow-Origin"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},
		Debug: true,
	})

	port := "5000"
	if port == "" {
		fmt.Println("default")
		port = "5000"
	}

	fmt.Println("todo listo en el", port, "!!")
	http.ListenAndServe(":"+port, c.Handler(router))

	return router

}
