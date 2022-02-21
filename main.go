package main

import (
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/Hassan-Zbib/GolangCloneApi/handlers"
)

func main() {
	
	router := mux.NewRouter()

	// users rountes
	router.HandleFunc("/api/users/friendlist", h.Friendlist).Methods("POST")
	router.HandleFunc("/api/users/addfriend", h.AddFriend).Methods("POST")
	router.HandleFunc("/api/users/acceptfriend", h.Accept).Methods("POST")

	// statuses rountes
	router.HandleFunc("/api/statuses/post", h.Post).Methods("POST")
	router.HandleFunc("/api/statuses/getfeed", h.GetFeed).Methods("POST")

	
	log.Println("Server started on: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}