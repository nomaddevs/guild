package recruits

import (
	"github.com/munsy/guild/config"
)

type Recruiting struct {
	DeathKnight bool
	DemonHunter bool
	Druid       bool
	Hunter      bool
	Mage        bool
	Monk        bool
	Paladin     bool
	Priest      bool
	Rogue       bool
	Shaman      bool
	Warlock     bool
	Warrior     bool
}

func (r *Recruiting) Set() error {
	db := &database.MariaDB{
		Username:       config.DBUsername,
		Unixsocketpath: config.DBUnixsocketpath,
		Password:       config.DBPassword,
		Host:           config.DBHost,
		Port:           config.DBPort,
		Database:       config.DBName,
		Charset:        config.DBCharset,
	}

	return db.SetRecruiting(r.DeathKnight, r.DemonHunter,
		r.Druid, r.Hunter, r.Mage, r.Monk, r.Paladin,
		r.Priest, r.Rogue, r.Shaman, r.Warlock, r.Warrior)
}
