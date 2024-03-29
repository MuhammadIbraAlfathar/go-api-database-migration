package app

import (
	"database/sql"
	"github.com/MuhammadIbraAlfathar/go-restful-api/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/golang_database_migration")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db

	// migrate -database "mysql://root@tcp(localhost:3306)/golang_database_migration" -path db/migration up
	// migrate -database "mysql://root@tcp(localhost:3306)/golang_database_migration" -path db/migration up 2
	// migrate -database "mysql://root@tcp(localhost:3306)/golang_database_migration" -path db/migration down
	// migrate -database "mysql://root@tcp(localhost:3306)/golang_database_migration" -path db/migration down 2
}
