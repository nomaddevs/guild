package database

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MariaDB struct {
	Username       string
	Unixsocketpath string
	Password       string
	Host           string
	Port           string
	Database       string
	Charset        string
}

func (db *MariaDB) DriverName() string {
	return "mysql"
}

// ConnectionString returns the connection string. Possible formats:
//  - user@unix(/path/to/socket)/dbname?charset=utf8
//  - user:password@/dbname
//  - user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname
//  - user:password@tcp(localhost:5555)/dbname?charset=utf8
func (db *MariaDB) ConnectionString() string {
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

func (db *MariaDB) Test() error {
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

/*
SQL Queries

CREATE DATABASE guild;

USE guild;

CREATE TABLE applications(id BIGINT NOT NULL AUTO_INCREMENT, battleid BIGINT NOT NULL, battletag varchar(50) NOT NULL,wowcharacter varchar(50) NOT NULL, email varchar(50) NOT NULL, realname varchar(50) NOT NULL, location varchar(100) NOT NULL, age TINYINT NOT NULL, gender varchar(20) NOT NULL, computerspecs varchar(500) NOT NULL, previousguilds varchar(500) NOT NULL, reasonsleavingguilds varchar(500) NOT NULL, whyjointhisguild varchar(500) NOT NULL, wowreferences varchar(500) NOT NULL, finalremarks varchar(500) NOT NULL, PRIMARY KEY (id)) ENGINE = InnoDB;

CREATE TABLE newsposts(id BIGINT NOT NULL AUTO_INCREMENT, title VARCHAR(128) NOT NULL, body MEDIUMBLOB NOT NULL, date DATETIME NOT NULL, author VARCHAR(64) NOT NULL, PRIMARY KEY (id)) ENGINE = InnoDB;

CREATE TABLE newscomments(id BIGINT NOT NULL AUTO_INCREMENT, pid BIGINT NOT NULL, title VARCHAR(128) NOT NULL, body MEDIUMBLOB NOT NULL, date DATETIME NOT NULL, author VARCHAR(64) NOT NULL, PRIMARY KEY (id), CONSTRAINT `fk_parent_post` FOREIGN KEY (pid) REFERENCES newsposts (id) ON DELETE CASCADE ON UPDATE RESTRICT) ENGINE = InnoDB;

INSERT INTO newsposts(title, body, date, author) values ("test 1", "here's some content", "2013-07-18 13:44:22.123456", "Munsy");

CREATE USER 'guild'@'localhost' IDENTIFIED BY 'a';

GRANT ALL PRIVILEGES ON *.* TO 'guild'@'localhost';
*/
