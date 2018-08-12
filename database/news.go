package database

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func (db *MariaDB) WriteNewsPost(title, body, author string) error {
	conn, err := sql.Open(db.DriverName(), db.ConnectionString())

	if nil != err {
		return err
	}
	defer conn.Close()

	in, err := conn.Prepare("INSERT INTO newsposts(title, body, date, author) values (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer in.Close()

	in.Exec(title, body, time.Now(), author)

	return nil
}

func (db *MariaDB) ReadNewsPosts() ([]int, []string, []string, []time.Time, []string, error) {
	var (
		id     int
		title  string
		body   string
		date   time.Time
		author string

		ids     []int
		titles  []string
		bodys   []string
		dates   []time.Time
		authors []string
	)

	conn, err := sql.Open(db.DriverName(), db.ConnectionString())

	if nil != err {
		return nil, nil, nil, nil, nil, err
	}
	defer conn.Close()

	rows, err := conn.Query("SELECT * from newsposts")

	if err != nil {
		return nil, nil, nil, nil, nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &title, &body, &date, &author)

		if err != nil {
			return nil, nil, nil, nil, nil, err
		}

		ids = append(ids, id)
		titles = append(titles, title)
		bodys = append(bodys, body)
		dates = append(dates, date)
		authors = append(authors, author)
	}

	err = rows.Err()

	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	return ids, titles, bodys, dates, authors, nil
}
