//go:build tools
//+build tools
package tools

import (
	_ "github.com/pressly/goose/v3/cmd/goose"
	_ "github.com/go-jet/jet/v2/cmd/jet"

)

//go run github.com/pressly/goose/v3/cmd/goose postgres "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" up

//go run github.com/go-jet/jet/v2/cmd/jet -dsn="postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" -path=./database/gen
