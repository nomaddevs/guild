package database

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func (db *MariaDB) Install() error {
	fmt.Printf("Creating table '%s.Applications...\n", db.Database)
	err := db.createTableApplications()

	if nil != err {
		return err
	}
	fmt.Printf("Successfully created table '%s.Applications\n", db.Database)

	fmt.Printf("Creating table '%s.NewsPosts...\n", db.Database)
	err = db.createTableNewsPosts()

	if nil != err {
		return err
	}
	fmt.Printf("Successfully created table '%s.NewsPosts\n", db.Database)

	return nil
}

func (db *MariaDB) createTableApplications() error {
	conn, err := sql.Open(db.DriverName(), db.ConnectionString())

	if nil != err {
		return err
	}
	defer conn.Close()

	statement := `CREATE TABLE applications(
		id BIGINT NOT NULL AUTO_INCREMENT,
		status INT NOT NULL,
		battleid BIGINT NOT NULL,
		battletag varchar(50) NOT NULL,
		wowcharacter varchar(50) NOT NULL,
		email varchar(50) NOT NULL, 
		realname varchar(50) NOT NULL, 
		location varchar(100) NOT NULL, 
		age TINYINT NOT NULL, 
		gender varchar(20) NOT NULL, 
		computerspecs varchar(500) NOT NULL, 
		previousguilds varchar(500) NOT NULL, 
		reasonsleavingguilds varchar(500) NOT NULL, 
		whyjointhisguild varchar(500) NOT NULL, 
		wowreferences varchar(500) NOT NULL, 
		finalremarks varchar(500) NOT NULL, 
		PRIMARY KEY (id)
		) ENGINE = InnoDB;`

	in, err := conn.Prepare(statement)
	if err != nil {
		return err
	}
	defer in.Close()

	in.Exec()

	return nil
}

func (db *MariaDB) createTableNewsPosts() error {
	conn, err := sql.Open(db.DriverName(), db.ConnectionString())

	if nil != err {
		return err
	}
	defer conn.Close()

	statement := `CREATE TABLE newsposts(
		id BIGINT NOT NULL AUTO_INCREMENT, 
		title VARCHAR(128) NOT NULL, 
		body MEDIUMBLOB NOT NULL, 
		date DATETIME NOT NULL, 
		author VARCHAR(64) NOT NULL, 
		PRIMARY KEY (id)
		) ENGINE = InnoDB;`

	in, err := conn.Prepare(statement)
	if err != nil {
		return err
	}
	defer in.Close()

	in.Exec()

	return nil
}

func (db *MariaDB) createTableNewsPostComments() error {
	conn, err := sql.Open(db.DriverName(), db.ConnectionString())

	if nil != err {
		return err
	}
	defer conn.Close()

	statement := `CREATE TABLE newscomments(
		id BIGINT NOT NULL AUTO_INCREMENT, 
		pid BIGINT NOT NULL, 
		title VARCHAR(128) NOT NULL, 
		body MEDIUMBLOB NOT NULL, 
		date DATETIME NOT NULL, 
		author VARCHAR(64) NOT NULL, 
		PRIMARY KEY (id), 
		CONSTRAINT ` + "`fk_parent_post`" + ` 
		FOREIGN KEY (pid) REFERENCES newsposts (id) 
		ON DELETE CASCADE ON UPDATE RESTRICT
		) ENGINE = InnoDB;`

	in, err := conn.Prepare(statement)
	if err != nil {
		return err
	}
	defer in.Close()

	in.Exec()

	return nil
}
