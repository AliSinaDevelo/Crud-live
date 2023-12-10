package main

import (
	"fmt"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"github.com/lib/pq"

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

	
}