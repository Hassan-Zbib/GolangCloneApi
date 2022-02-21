package h

import (
	"encoding/json"
	"net/http"
	"reflect"
	_ "github.com/go-sql-driver/mysql"
	"github.com/Hassan-Zbib/GolangCloneApi/config"
)

// structs - classes 
type friends struct {
	ID        int         `json:"id"`
	Name      string      `json:"name"`
	Email     string      `json:"email"`
	Password  string      `json:"password"`
	Picture   interface{} `json:"picture"`
	CreatedAt string      `json:"created_at"`
	UpdatedAt string      `json:"updated_at"`
	RequestID int         `json:"request_id"`
	Request   string      `json:"request"`
	AddedAt   string      `json:"added_at"`
}

type feed struct {
	ID         int    `json:"id"`
	UserID     int    `json:"user_id"`
	Content    string `json:"content"`
	LikesCount int    `json:"likes_count"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	IsLiked    bool   `json:"is_liked"`
}

type request struct {
	UserID int `json:"user_id"`
	RecordID int `json:"record_id"`
	Content int `json:"content"`
}
type response struct {
	Message string `json:"message"`
}


// users
func AddFriend(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// db connect
	db := config.DbConn()
	// get json request
	decoder := json.NewDecoder(r.Body)

	var vars request
	err := decoder.Decode(&vars)
	if err != nil {
        panic(err)
    }

	user_id := vars.UserID
	friend_id := vars.RecordID
	request := "pending"

	// perform a db.Query
	stmt, err := db.Prepare("INSERT INTO friends(user_id,friend_id,request) VALUES (?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(user_id, friend_id, request)
	if err != nil {
		panic(err.Error())
	}
	res := response{
		Message: "Request Sent",
	}
	json.NewEncoder(w).Encode(res)
	defer db.Close()
}

func Accept(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// db connect
	db := config.DbConn()
	// get json request
	decoder := json.NewDecoder(r.Body)

	var vars request
	err := decoder.Decode(&vars)
	if err != nil {
        panic(err)
    }

	record_id := vars.RecordID
	request := "accepted"

	// perform a db.Query
	stmt, err := db.Prepare("UPDATE friends SET request=? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(request, record_id)
	if err != nil {
		panic(err.Error())
	}
	res := response{
		Message: "Accepted",
	}
	json.NewEncoder(w).Encode(res)
	defer db.Close()
}

func Friendlist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// db connect
	db := config.DbConn()
	// get json request
	decoder := json.NewDecoder(r.Body)

	var vars request
	err := decoder.Decode(&vars)
	if err != nil {
        panic(err)
    }

	user_id := vars.UserID
	request := "accepted"

	var list []friends

	result, err := db.Query(`SELECT u.*, f.id as request_id , f.request, f.created_at as added_at 
								FROM users u 
								INNER JOIN friends f  
								ON  f.user_id = u.id OR f.friend_id = u.id  
								WHERE u.id != ? AND f.request = ?`, user_id, request)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		var record friends

        s := reflect.ValueOf(&record).Elem()
        numCols := s.NumField()
        columns := make([]interface{}, numCols)
        for i := 0; i < numCols; i++ {
            field := s.Field(i)
            columns[i] = field.Addr().Interface()
        }

		err = result.Scan(columns...)
		if err != nil {
			panic(err.Error())
		}

		list = append(list, record)
	}
	json.NewEncoder(w).Encode(list)
	defer db.Close()
}



// statuses
func Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// db connect
	db := config.DbConn()
	// get json request
	decoder := json.NewDecoder(r.Body)

	var vars request
	err := decoder.Decode(&vars)
	if err != nil {
        panic(err)
    }

	user_id := vars.UserID
	content := vars.Content

	// perform a db.Query
	stmt, err := db.Prepare("INSERT INTO statuses (user_id, content) VALUES (?, ?);")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(user_id, content)
	if err != nil {
		panic(err.Error())
	}
	res := response{
		Message: "Status Created",
	}
	json.NewEncoder(w).Encode(res)
	defer db.Close()
}

func GetFeed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// db connect
	db := config.DbConn()
	// get json request
	decoder := json.NewDecoder(r.Body)

	var vars request
	err := decoder.Decode(&vars)
	if err != nil {
        panic(err)
    }

	user_id := vars.UserID
	request := "accepted"

	var list []feed

	result, err := db.Query(`SELECT s.*, u.name, u.email,  CASE 
								WHEN ? in (SELECT DISTINCT user_id FROM likes WHERE status_id = s.id) THEN 1
								ELSE 0
							END as is_liked
							FROM statuses s
							INNER JOIN users u ON s.user_id = u.id 
							INNER JOIN friends f ON s.user_id = f.user_id OR s.user_id = f.friend_id
							WHERE s.user_id != ? AND f.request = ?
							;`, user_id, user_id, request)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		var record feed

        s := reflect.ValueOf(&record).Elem()
        numCols := s.NumField()
        columns := make([]interface{}, numCols)
        for i := 0; i < numCols; i++ {
            field := s.Field(i)
            columns[i] = field.Addr().Interface()
        }

		err = result.Scan(columns...)
		if err != nil {
			panic(err.Error())
		}

		list = append(list, record)
	}
	json.NewEncoder(w).Encode(list)
	defer db.Close()
}