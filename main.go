package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type User struct {
	ID	int	`json:"id"`
	NAME string `json:"name"`
	EMAIL string `json:"mail"`
}

func main() {
	// Connect to Postgres Instance
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// create the table if doesn't exist
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT, email TEXT)")
	if err != nil {
		log.Fatal(err)
	}

	// create router
	router := mux.NewRouter()
	router.HandleFunc("/users", getUsers(db)).Methods("GET")
	router.HandleFunc("/users/{id}", getUsers(db)).Methods("GET")
	router.HandleFunc("/users", createUser(db)).Methods("POST")
	router.HandleFunc("/users/{id}", updateUser(db)).Methods("PUT")
	router.HandleFunc("/users/{id}", deleteUser(db)).Methods("DELETE")

	// start server
	log.Fatal(http.ListenAndServe(":8000", jsonContTMiddleWare(router)))
}

func jsonContTMiddleWare (next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "applications/json")
		next.ServeHTTP(w, r)
	})
}

// get all the users
func getUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r* http.Request) {
		rows, err := db.Query("SELECT * FROM users")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		user := []User{}
		for rows.Next() {
			var u User
			if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
				log.Fatal(err)
			}
			user = append(user, u)
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(user)
	}
}

// get user by ID
func getUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r* http.Request) {
		params := mux.Vars(r)
		row := db.QueryRow("SELECT * FROM users WHERE id=$1", params["id"])
		var u User
		if err := row.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(u)
	}
}

// create user
func createUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r* http.Request) {
		var u User
		json.NewDecoder(r.Body).Decode(&u)
		_, err := db.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", u.Name, u.Email)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(u)
	}
}
// update user
func updateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r* http.Request) {
		params := mux.Vars(r)

		var u User
		json.NewDecoder(r.Body).Decode(&u)
		_, err := db.Exec("UPDATE users SET name=$1, email=$2 WHERE id=$3", u.Name, u.Email, params["id"])
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(u)
	}
}

// delete user
func deleteUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r* http.Request) {
		params := mux.Vars(r)
		_, err := db.Exec("DELETE FROM users WHERE id=$1", params["id"])
		if err != nil {
			// todo fix err handle
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(params["id"])
	}
}