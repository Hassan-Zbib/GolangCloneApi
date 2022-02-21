package h

import (
	"encoding/json"
	"fmt"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/Hassan-Zbib/GolangCloneApi/config"
)

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
	Number  string `json:"number"`
}


func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := config.DbConn()

	var users []User

	result, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		var post User
		err = result.Scan(&post.ID, &post.Name, &post.Country, &post.Number)
		if err != nil {
			panic(err.Error())
		}

		users = append(users, post)
	}
	json.NewEncoder(w).Encode(users)
	defer db.Close()
}
func InsertUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := config.DbConn()
	vars := mux.Vars(r)
	Name := vars["name"]
	Country := vars["country"]
	Number := vars["number"]

	// perform a db.Query insert
	stmt, err := db.Prepare("INSERT INTO users(name, country, number) VALUES(?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(Name, Country, Number)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "New user was created")
	defer db.Close()
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := config.DbConn()
	params := mux.Vars(r)

	// perform a db.Query insert
	stmt, err := db.Query("SELECT * FROM users WHERE id = ?", params["id"])
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	var post User
	for stmt.Next() {

		err = stmt.Scan(&post.ID, &post.Name, &post.Country, &post.Number)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(post)
	defer db.Close()
}
func DelUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := config.DbConn()
	params := mux.Vars(r)

	// perform a db.Query insert
	stmt, err := db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(params["id"])
	fmt.Fprintf(w, "User with ID = %s was deleted", params["id"])
	defer db.Close()
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := config.DbConn()
	params := mux.Vars(r)
	Name := params["name"]
	Country := params["country"]
	Number := params["number"]

	// perform a db.Query insert
	stmt, err := db.Prepare("Update users SET name = ?, country = ?, number = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(Name, Country, Number, params["id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "User with ID = %s was updated", params["id"])
	defer db.Close()
}