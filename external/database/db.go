package database

import (
	"api/external/logger"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB(l *logger.Logger) *sql.DB {
	db, err := sql.Open("mysql",
		"user:password@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		l.Errorf(err.Error())
	}
	return db
}
