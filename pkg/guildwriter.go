package models

import(
        "database/sql"
        "errors"
        "io/ioutil"
        "encoding/json"
        "time"
)

type GuildWriter struct {
	Settings *MariaDBConfig
}

func NewGuildWriter() (*GuildWriter, error) {
        db := &MariaDBConfig{}

        filename := "mariadbconfig.json"

        data, err := ioutil.ReadFile(filename)
        if nil != err {
                return nil, err
        }

        err = json.Unmarshal(data, &db)
        if nil != err {
                return nil, err
        }

        gw := &GuildWriter{
                db,
        }

        return gw, nil
}

func (gw *GuildWriter) RefreshRoster() error {
	conn, err := sql.Open(gw.Settings.DriverName(), gw.Settings.ConnectionString())
        if nil != err {
                panic(err)
        }
        defer conn.Close()

        var key string
        conn.QueryRow("SELECT apikey FROM bnetapi").Scan(&key)

        if len(key) != 32 {
                return errors.New("Bad length: " + string(len(key)))
        }

	var Guildinfo GuildInfo
        url := "https://us.api.battle.net/wow/guild/thrall/NoBelfsAllowed?fields=members&locale=en_US&apikey=" + key
        err = Get(url, &Guildinfo)
        if nil != err {
                return err
        }

	insert := "INSERT INTO roster(name, rank, class, race, gender, level," 
	insert += "acheivementpoints, thumbnail, lastmodified) VALUES(?, ?, ?, ?, ?, ?, ?)"

        stmtIns, err := conn.Prepare(insert)
        if err != nil {
                return err
        }
        defer stmtIns.Close()

	for _, v := range Guildinfo.Members {
		_, err = stmtIns.Exec(v.Character.Name,
					v.Rank,
					v.Character.Class,
					v.Character.Race,
					v.Character.Gender,
					v.Character.Level,
					v.Character.AchievementPoints,
					v.Character.Thumbnail,
					v.Character.LastModified)
		if err != nil {
			return err
		}
	}

	return nil
}

func (gw *GuildWriter) WriteNewsPost(np *NewsPost, user BnetUser) error {
        // Create the database handle, confirm driver is present
        conn, err := sql.Open(gw.Settings.DriverName(), gw.Settings.ConnectionString())
        if nil != err {
                return err
        }
        defer conn.Close()

        err = conn.Ping()
        if err != nil {
                return err
        }

        stmtIns, err := conn.Prepare("INSERT INTO newsposts VALUES( ?, ?, ?, ?, ? )") // ? = placeholder
        if err != nil {
                return err
        }
        defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

        var name string
        if  "" != user.BattleTag {
                name = user.BattleTag
        } else {
                name = "Anonymous"
        }

        postTime := time.Now().Format(time.RFC3339)

        _, err = stmtIns.Exec(np.ID, np.Title, np.Body, postTime, name) // Insert tuples (i, i^2)
        if err != nil {
                return err
        }

        return nil
}

func (gw *GuildWriter) CreateApplicant(app *AppInfo) error {
        // Create the database handle, confirm driver is present
        conn, err := sql.Open(gw.Settings.DriverName(), gw.Settings.ConnectionString())
        if nil != err {
                return err
        }
        defer conn.Close()

        err = conn.Ping()
        if err != nil {
                return err
        }

        statement := "INSERT INTO applications(wowcharacter, email,"
        statement += "realname,location,age,gender,computerspecs,previousguilds,"
        statement += "reasonsleavingguilds,whyjointhisguild,wowreferences,finalremarks) "
        statement += "VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

        stmtIns, err := conn.Prepare(statement)
        if err != nil {
                return err
        }
        defer stmtIns.Close() // Close the statement when we leave main() / the program terminates


        _, err = stmtIns.Exec(app.Character,
                        app.Email,
                        app.RealName,
                        app.Location,
                        app.Age,
                        app.Gender,
                        app.ComputerSpecs,
                        app.PreviousGuilds,
                        app.ReasonsLeavingGuilds,
                        app.WhyJoinThisGuild,
                        app.References,
                        app.FinalRemarks)

        if err != nil {
                return err
        }

        return nil
}
