package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB(connectionsString string) error {
	var err error
	DB, err = sql.Open("mysql", connectionsString)
	if err != nil {
		return err
	}
	return DB.Ping()
}
