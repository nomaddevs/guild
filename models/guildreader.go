package models

import(
	"database/sql"
	"errors"
	"io/ioutil"
	"encoding/json"
//	"time"
)

type GuildReader struct {
        Settings *MariaDBConfig
}

func NewGuildReader() (*GuildReader, error) {
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

	gr := &GuildReader{
		db,
	}

	return gr, nil
}

func (gr *GuildReader) GetBNetAPIKey() (string, error) {
        conn, err := sql.Open(gr.Settings.DriverName(), gr.Settings.ConnectionString())
        if nil != err {
                panic(err)
        }
        defer conn.Close()

        var key string
        conn.QueryRow("SELECT apikey FROM bnetapi").Scan(&key)

        if len(key) != 32 {
                return "", errors.New("Bad length: " + string(len(key)))
        }

        return key, nil
}

func (gr *GuildReader) GetNewsPosts() ([]NewsPost, error) {
        // Create the database handle, confirm driver is present
        conn, err := sql.Open(gr.Settings.DriverName(), gr.Settings.ConnectionString())
        if nil != err {
                return nil, err
        }
        defer conn.Close()

        err = conn.Ping()
        if err != nil {
                return nil, err
        }

        // Execute the query
        rows, err := conn.Query("SELECT * FROM newsposts ORDER BY date")
        if err != nil {
                return nil, err
        }

        npList := []NewsPost{}

        // Fetch rows
        for rows.Next() {
                var np NewsPost
                err = rows.Scan(&np.ID, &np.Title, &np.Body, &np.Date, &np.Author)
                if err != nil {
                        return nil, err
                }

                npList = append(npList, np)
        }

        return npList, nil
}

