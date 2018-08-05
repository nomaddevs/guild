package models

import (
	"github.com/munsy/guild/config"
	"github.com/munsy/guild/database"
)

type User struct {
	ID        int
	BattleTag string
	Applied   bool
}

func Applied(id int) (bool, error) {
	db := &database.MariaDB{
		Username:       config.DBUsername,
		Unixsocketpath: config.DBUnixsocketpath,
		Password:       config.DBPassword,
		Host:           config.DBHost,
		Port:           config.DBPort,
		Database:       config.DBName,
		Charset:        config.DBCharset,
	}

	return db.GetApplicant(id)
}
