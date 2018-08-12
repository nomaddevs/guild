package recruits

import (
	"github.com/munsy/guild/config"
	"github.com/munsy/guild/database"
)

type DeathKnight struct {
	Desired int
	Filled  int
	Pending int
}

func (c DeathKnight) Need() bool {
	return c.Desired-c.Filled != 0
}

type DemonHunter struct {
	Desired int
	Filled  int
	Pending int
}

func (c DemonHunter) Need() bool {
	return c.Desired-c.Filled != 0
}

type Druid struct {
	Desired int
	Filled  int
	Pending int
}

func (c Druid) Need() bool {
	return c.Desired-c.Filled != 0
}

type Hunter struct {
	Desired int
	Filled  int
	Pending int
}

func (c Hunter) Need() bool {
	return c.Desired-c.Filled != 0
}

type Mage struct {
	Desired int
	Filled  int
	Pending int
}

func (c Mage) Need() bool {
	return c.Desired-c.Filled != 0
}

type Monk struct {
	Desired int
	Filled  int
	Pending int
}

func (c Monk) Need() bool {
	return c.Desired-c.Filled != 0
}

type Paladin struct {
	Desired int
	Filled  int
	Pending int
}

func (c Paladin) Need() bool {
	return c.Desired-c.Filled != 0
}

type Priest struct {
	Desired int
	Filled  int
	Pending int
}

func (c Priest) Need() bool {
	return c.Desired-c.Filled != 0
}

type Rogue struct {
	Desired int
	Filled  int
	Pending int
}

func (c Rogue) Need() bool {
	return c.Desired-c.Filled != 0
}

type Shaman struct {
	Desired int
	Filled  int
	Pending int
}

func (c Shaman) Need() bool {
	return c.Desired-c.Filled != 0
}

type Warlock struct {
	Desired int
	Filled  int
	Pending int
}

func (c Warlock) Need() bool {
	return c.Desired-c.Filled != 0
}

type Warrior struct {
	Desired int
	Filled  int
	Pending int
}

func (c Warrior) Need() bool {
	return c.Desired-c.Filled != 0
}

type Recruiting struct {
	DeathKnight DeathKnight
	DemonHunter DemonHunter
	Druid       Druid
	Hunter      Hunter
	Mage        Mage
	Monk        Monk
	Paladin     Paladin
	Priest      Priest
	Rogue       Rogue
	Shaman      Shaman
	Warlock     Warlock
	Warrior     Warrior
}

func Get() *Recruiting {
	db := &database.MariaDB{
		Username:       config.DBUsername,
		Unixsocketpath: config.DBUnixsocketpath,
		Password:       config.DBPassword,
		Host:           config.DBHost,
		Port:           config.DBPort,
		Database:       config.DBName,
		Charset:        config.DBCharset,
	}

	db.GetRecruiting()

	return nil
}

func Set(r *Recruiting) error {
	db := &database.MariaDB{
		Username:       config.DBUsername,
		Unixsocketpath: config.DBUnixsocketpath,
		Password:       config.DBPassword,
		Host:           config.DBHost,
		Port:           config.DBPort,
		Database:       config.DBName,
		Charset:        config.DBCharset,
	}

	db.SetRecruiting(r.DeathKnight.Desired, r.DeathKnight.Filled, r.DeathKnight.Pending)
	db.SetRecruiting(r.DemonHunter.Desired, r.DemonHunter.Filled, r.DemonHunter.Pending)
	db.SetRecruiting(r.Druid.Desired, r.Druid.Filled, r.Druid.Pending)
	db.SetRecruiting(r.Hunter.Desired, r.Hunter.Filled, r.Hunter.Pending)
	db.SetRecruiting(r.Mage.Desired, r.Mage.Filled, r.Mage.Pending)
	db.SetRecruiting(r.Monk.Desired, r.Monk.Filled, r.Monk.Pending)
	db.SetRecruiting(r.Paladin.Desired, r.Paladin.Filled, r.Paladin.Pending)
	db.SetRecruiting(r.Priest.Desired, r.Priest.Filled, r.Priest.Pending)
	db.SetRecruiting(r.Rogue.Desired, r.Rogue.Filled, r.Rogue.Pending)
	db.SetRecruiting(r.Shaman.Desired, r.Shaman.Filled, r.Shaman.Pending)
	db.SetRecruiting(r.Warlock.Desired, r.Warlock.Filled, r.Warlock.Pending)
	db.SetRecruiting(r.Warrior.Desired, r.Warrior.Filled, r.Warrior.Pending)

	return nil
}