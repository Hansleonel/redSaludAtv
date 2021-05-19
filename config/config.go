package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var password string = "developer2020RS#$1"
var driver string = "mysql"
var user string = "rsdevus1"
var host string = "rdsdevenv.czvvckkesgis.us-east-2.rds.amazonaws.com"
var nombre2 string = "global_rs"

func GetMySQLDB() (db *sql.DB, err error) {
	db, err = sql.Open(driver, user+":"+password+"@tcp("+host+")/"+nombre2)
	db.SetMaxOpenConns(50)
	return
}
