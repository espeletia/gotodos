package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"todos/database/gen/postgres/public/model"
	"todos/database/gen/postgres/public/table"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/gorilla/mux"
)

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func CreateUser(dbconn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := User{}
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		stmt := table.Users.INSERT(table.Users.Name).VALUES(user.Name).RETURNING(table.Users.AllColumns)
		fmt.Printf("stmt: %s\n", stmt.DebugSql())
		dest := model.Users{}
		err = stmt.Query(dbconn, &dest)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		user.ID = int64(dest.ID)
		
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}
}

func GetUser(dbconn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		stmt := table.Users.SELECT(table.Users.AllColumns).WHERE(table.Users.ID.EQ(postgres.Int(int64(id))))
		var dest model.Users
		err = stmt.Query(dbconn, &dest)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		user := User{ID: int64(dest.ID), Name: dest.Name}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	}
}
