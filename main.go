package main

import (
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/Hassan-Zbib/GolangCloneApi/handlers"
)


func main() {
	log.Println("Server started on: http://localhost:8080")
	router := mux.NewRouter()

	// users rountes
	router.HandleFunc("/all", h.Index).Methods("GET")
	router.HandleFunc("/add", h.InsertUser).Methods("POST").Queries("name", "{name}", "country", "{country}", "number", "{number}")
	router.HandleFunc("/get/{id}", h.GetUser).Methods("GET")

	// statuses rountes
	router.HandleFunc("/update/{id}", h.UpdateUser).Methods("PUT").Queries("name", "{name}", "country", "{country}", "number", "{number}")
	router.HandleFunc("/del/{id}", h.DelUser).Methods("DELETE")

	http.ListenAndServe(":8080", router)

}