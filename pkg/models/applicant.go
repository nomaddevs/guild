package models

import (
	"github.com/munsy/guild/database"
)

type Applicant struct {
	Battletag            string
	Character            string
	Email                string
	RealName             string
	Location             string
	Age                  string
	Gender               string
	ComputerSpecs        string
	PreviousGuilds       string
	ReasonsLeavingGuilds string
	WhyJoinThisGuild     string
	References           string
	FinalRemarks         string
}

func (a *Applicant) Save() error {
	db := &database.MariaDB{
		Username:       config.DBUsername,
		Unixsocketpath: config.DBUnixsocketpath,
		Password:       config.DBPassword,
		Host:           config.DBHost,
		Port:           config.DBPort,
		Database:       config.DBName,
		Charset:        config.DBCharset,
	}

	return db.WriteApplicant(a)
}
