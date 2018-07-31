package config

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

type DBconfig interface {
	DriverName() string       // "mysql", "mssql", etc
	ConnectionString() string // username:password@tcp(host:port)/database for mysql
	Test() error              // See https://github.com/golang/go/wiki/SQLDrivers for a list of SQL drivers.
}

type dbquery struct {
	query string
}
