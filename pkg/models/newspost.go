package models

import (
	"time"

	"github.com/munsy/guild/config"
	"github.com/munsy/guild/database"
)

type NewsPost struct {
	ID     int
	Title  string
	Body   string
	Date   time.Time
	Author string
}

func (np *NewsPost) Save() error {
	db := &database.MariaDB{
		Username:       config.DBUsername,
		Unixsocketpath: config.DBUnixsocketpath,
		Password:       config.DBPassword,
		Host:           config.DBHost,
		Port:           config.DBPort,
		Database:       config.DBName,
		Charset:        config.DBCharset,
	}

	return db.WriteNewsPost(np.Title, np.Body, np.Author)
}

type NewsPosts []NewsPost

func (nps NewsPosts) Read() error {
	db := &database.MariaDB{
		Username:       config.DBUsername,
		Unixsocketpath: config.DBUnixsocketpath,
		Password:       config.DBPassword,
		Host:           config.DBHost,
		Port:           config.DBPort,
		Database:       config.DBName,
		Charset:        config.DBCharset,
	}

	ids, titles, bodies, dates, authors, err := db.ReadNewsPosts()

	if nil != err {
		return err
	}

	for i := 0; i < len(ids); i++ {
		np := NewsPost{
			ID:     ids[i],
			Title:  titles[i],
			Body:   bodies[i],
			Date:   dates[i],
			Author: authors[i],
		}

		nps = append(nps, np)
	}

	return nil
}
