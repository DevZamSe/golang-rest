package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetMySQLDB() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbUser := "bode_test"
	dbPass := "Test123456$"
	dbName := "bode_pacha"
	dbHost := "mysql-bode.alwaysdata.net"
	// db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbHost+")/"+dbName)
	return
}
