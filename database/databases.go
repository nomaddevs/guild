package conf

import (
	"database/sql"
)

type DBconfig interface {
	DriverName() string          // "mysql", "mssql", etc
	ConnectionString() string    // name:password@/database for mysql
	Test() error                 // See https://github.com/golang/go/wiki/SQLDrivers for a list of SQL drivers.
	GetTLS() (*TLSconfig, error) // Gets the TLS info from table
}

type dbquery struct {
	query string
}

// change this later
//func GetAPICredential(column string) string {
func GetAPICredential(username, password, host, port, database, table, column string) string {
	//db, err := sql.Open("mysql", "<username>:<pw>@tcp(<HOST>:<port>)/<dbname>")
	db, err := sql.Open("mysql", username+":"+password+"@tcp("+host+":"+port+")/"+database)
	if nil != err {
		panic(err)
	}
	defer db.Close()

	var key string
	db.QueryRow("SELECT " + column + " FROM " + table).Scan(&key)

	if len(key) != 32 {
		panic("Bad API credential length from column: " + column + " in table: " + table + "\ntrying to find key: " + key + "with length: " + string(len(key)))
	}

	return key
}

/*
===================================================================
|Some of The Most Important SQL Commands                          |
|-----------------------------------------------------------------|
|SELECT             - extracts data from a database               |
|UPDATE             - updates data in a database                  |
|DELETE             - deletes data from a database                |
|INSERT INTO        - inserts new data into a database            |
|CREATE DATABASE    - creates a new database                      |
|ALTER DATABASE     - modifies a database                         |
|CREATE TABLE       - creates a new table                         |
|ALTER TABLE        - modifies a table                            |
|DROP TABLE         - deletes a table                             |
|CREATE INDEX       - creates an index (search key)               |
|DROP INDEX         - deletes an index                            |
-------------------------------------------------------------------
*/
