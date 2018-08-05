package models

import (
	"github.com/munsy/guild/config"
	"github.com/munsy/guild/database"
)

type Applicant struct {
	BattleID             int
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

	return db.WriteApplicant(a.BattleID, a.Battletag, a.Character,
		a.Email, a.RealName, a.Location, a.Age, a.Gender,
		a.ComputerSpecs, a.PreviousGuilds, a.ReasonsLeavingGuilds,
		a.WhyJoinThisGuild, a.References, a.FinalRemarks)
}
