package main

import (
	"fmt"

	"github.com/XSAM/otelsql"
	_ "github.com/lib/pq"
)

func main() {
	dbconn, err := otelsql.Open("postgres", "postgres://postgres:postgres@localhost:5434/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}

	err = dbconn.Ping()
	if err != nil {
		panic(err)
	}
	
	dbconn.Query("INSERT INTO users (name) VALUES ('test')")
	rows, err := dbconn.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", rows)
}
