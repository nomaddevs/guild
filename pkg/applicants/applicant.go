package applicants

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

func View(id int) ([]Applicant, error) {
	db := &database.MariaDB{
		Username:       config.DBUsername,
		Unixsocketpath: config.DBUnixsocketpath,
		Password:       config.DBPassword,
		Host:           config.DBHost,
		Port:           config.DBPort,
		Database:       config.DBName,
		Charset:        config.DBCharset,
	}

	var apps []Applicant

	a, b, c, d, e, f, g, h, i, j, k, l, m, n, err := db.ViewApplicant(id)

	if nil != err {
		return nil, err
	}

	for x := 0; x < len(a); x++ {
		app := Applicant{
			BattleID:             a[x],
			Battletag:            b[x],
			Character:            c[x],
			Email:                d[x],
			RealName:             e[x],
			Location:             f[x],
			Age:                  g[x],
			Gender:               h[x],
			ComputerSpecs:        i[x],
			PreviousGuilds:       j[x],
			ReasonsLeavingGuilds: k[x],
			WhyJoinThisGuild:     l[x],
			References:           m[x],
			FinalRemarks:         n[x],
		}
		apps = append(apps, app)
	}

	return apps, nil
}

func ViewAll() ([]Applicant, error) {
	db := &database.MariaDB{
		Username:       config.DBUsername,
		Unixsocketpath: config.DBUnixsocketpath,
		Password:       config.DBPassword,
		Host:           config.DBHost,
		Port:           config.DBPort,
		Database:       config.DBName,
		Charset:        config.DBCharset,
	}

	var apps []Applicant

	a, b, c, d, e, f, g, h, i, j, k, l, m, n, err := db.ViewAllApplicants()

	if nil != err {
		return nil, err
	}

	for x := 0; x < len(a); x++ {
		app := Applicant{
			BattleID:             a[x],
			Battletag:            b[x],
			Character:            c[x],
			Email:                d[x],
			RealName:             e[x],
			Location:             f[x],
			Age:                  g[x],
			Gender:               h[x],
			ComputerSpecs:        i[x],
			PreviousGuilds:       j[x],
			ReasonsLeavingGuilds: k[x],
			WhyJoinThisGuild:     l[x],
			References:           m[x],
			FinalRemarks:         n[x],
		}
		apps = append(apps, app)
	}

	return apps, nil
}
