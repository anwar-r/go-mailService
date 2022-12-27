package router

import (
	"fmt"
	"log"
	ctr "mailService/Controller"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func HandleFunc() {
	router := mux.NewRouter()
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	}).Handler(router)
	router.HandleFunc("/mail", ctr.Mail).Methods("POST")

	fmt.Println("Go Mail Services with Go Clean Architecture")
	fmt.Println("Server started on: http://localhost:5050")

	log.Fatal(http.ListenAndServe(":5050", corsHandler))

}
