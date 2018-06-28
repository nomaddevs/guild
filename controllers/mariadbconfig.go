/*
SQL Queries

CREATE DATABASE guild;

USE guild;

CREATE TABLE tlsinfo(addr varchar(256) NOT NULL PRIMARY KEY, certfile varchar(256) NOT NULL, keyfile varchar(256) NOT NULL);

CREATE TABLE CREATE TABLE bnetapi(apikey varchar(32) NOT NULL PRIMARY KEY, apisecret varchar(32) NOT NULL);

CREATE TABLE newsposts(id BIGINT NOT NULL AUTO_INCREMENT, title VARCHAR(128) NOT NULL, body MEDIUMBLOB NOT NULL, date DATETIME NOT NULL, author VARCHAR(64) NOT NULL, PRIMARY KEY (id)) ENGINE = InnoDB;

CREATE TABLE newscomments(id BIGINT NOT NULL AUTO_INCREMENT, pid BIGINT NOT NULL, title VARCHAR(128) NOT NULL, body MEDIUMBLOB NOT NULL, date DATETIME NOT NULL, author VARCHAR(64) NOT NULL, PRIMARY KEY (id), CONSTRAINT `fk_parent_post` FOREIGN KEY (pid) REFERENCES newsposts (id) ON DELETE CASCADE ON UPDATE RESTRICT) ENGINE = InnoDB;

INSERT INTO newsposts(title, body, date, author) values ("test 1", "here's some content", "2013-07-18 13:44:22.123456", "Munsy"), ("test 2", "here's some more content", "2013-07-19 11:42:42.123006", "Munsy"), ("test 3", "here's some MORE content", "2013-08-21 6:12:24.263886", "Munsy");

INSERT INTO bnetapi(apikey, apisecret) VALUES ('', '');

CREATE USER 'guild'@'localhost' IDENTIFIED BY 'a';

GRANT ALL PRIVILEGES ON *.* TO 'guild'@'localhost';
*/

package controllers

import (
	"database/sql"
	"errors"

	"github.com/munsylol/guild/models"
	_ "github.com/go-sql-driver/mysql"
)

type MariaDBConfig struct {
	Username       string `json:"username"`
	Unixsocketpath string `json:"unixsocketpath"`
	Password       string `json:"password"`
	Host           string `json:"host"`
	Port           string `json:"port"`
	Database       string `json:"database"`
	Charset        string `json:"charset"`
}

// Public methods

func (db *MariaDBConfig) DriverName() string {
	return "mysql"
}

/* ConnectionString possible formats:
 * - [x] user@unix(/path/to/socket)/dbname?charset=utf8
 * - [x] user:password@/dbname
 * - [x] user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname
 * - [x] user:password@tcp(localhost:5555)/dbname?charset=utf8
 */
func (db *MariaDBConfig) ConnectionString() string {
	var connString string

	if "" == db.Username || "" == db.Database {
		return ""
	}
	if "" == db.Password && "" == db.Unixsocketpath {
		return ""
	}
	if "" != db.Unixsocketpath {
		connString = db.Username + "@unix(" + db.Unixsocketpath + ")/" + db.Database
		if "" != db.Charset {
			connString += "?charset=" + db.Charset
		}
		return connString
	}
	if "" == db.Host || "" == db.Port {
		connString = db.Username + ":" + db.Password + "@/" + db.Database
		if "" != db.Charset {
			connString += "?charset=" + db.Charset
		}
		return connString
	}

	connString = db.Username + ":" + db.Password + "@tcp(" + db.Host + ":" + db.Port + ")/" + db.Database
	if "" != db.Charset {
		connString += "?charset=" + db.Charset
	}

	return connString
}

func (db *MariaDBConfig) Test() error {
	// Create the database handle, confirm driver is present
	conn, err := sql.Open(db.DriverName(), db.ConnectionString())
	if nil != err {
		return err
	}
	defer conn.Close()

	err = conn.Ping()
	if err != nil {
		return err
	}

	version := ""
	conn.QueryRow("SELECT VERSION()").Scan(&version)
	if "" == version {
		return errors.New("Couldn't get SQL version.")
	}

	return nil
}

func (db *MariaDBConfig) GetTLS() (*TLSconfig, error) {
	// Create the database handle, confirm driver is present
	conn, err := sql.Open(db.DriverName(), db.ConnectionString())
	if nil != err {
		return nil, err
	}
	defer conn.Close()

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	addr := ""
	certfile := ""
	keyfile := ""

	conn.QueryRow("SELECT * FROM tlsinfo").Scan(&addr, &certfile, &keyfile)
	if "" == addr || "" == certfile || "" == keyfile {
		return nil, errors.New("Error retrieving TLS info.")
	}

	cfg := &TLSconfig{
		addr,
		certfile,
		keyfile,
	}

	return cfg, nil
}

func GetNewsPosts(username, password string) ([]models.NewsPost, error) {
	db := &MariaDBConfig{
                username,
                "",
                password,
                "localhost",
                "3306",
                "guild",
                "",
        }

        // Create the database handle, confirm driver is present
        conn, err := sql.Open(db.DriverName(), db.ConnectionString())
        if nil != err {
                return nil, err
        }
        defer conn.Close()

        err = conn.Ping()
        if err != nil {
                return nil, err
        }

	// Execute the query
	rows, err := conn.Query("SELECT * FROM newsposts")
	if err != nil {
		return nil, err
	}

	npList := []models.NewsPost{}

	// Fetch rows
	for rows.Next() {
		var np models.NewsPost
		err = rows.Scan(&np.ID, &np.Title, &np.Body, &np.Date, &np.Author)
		if err != nil {
			return nil, err
		}

		npList = append(npList, np)
	}

        return npList, nil
}

