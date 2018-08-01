package beta

import (
	"fmt"
	"net/http"

	"github.com/munsy/guild/config"
	"github.com/munsy/guild/database"
	"github.com/munsy/guild/pkg/models"
)

/*
type NewsPost struct {
	ID     int
	Title  string
	Body   string
	Date   time.Time
	Author string
}
*/

// News creates a single news post or returns a set of posts, depending on the http method.
func News(w http.ResponseWriter, r *http.Request) {
	db := &database.MariaDB{
		Username:       config.DBUsername,
		Unixsocketpath: config.DBUnixsocketpath,
		Password:       config.DBPassword,
		Host:           config.DBHost,
		Port:           config.DBPort,
		Database:       config.DBName,
		Charset:        config.DBCharset,
	}

	switch r.Method {
	case "GET":
		data, err := db.ReadNewsPosts()

		if nil != err {
			a.JSON(w, err)
		}

		a.JSON(w, data)

		break
	case "POST":
		title := r.FormValue("post_title")
		body := r.FormValue("post_body")
		author := r.FormValue("post_author")

		err := db.WriteNewsPost(title, body, author)

		if nil != err {
			a.JSON(w, err)
		}

		data, err := db.ReadNewsPosts()

		if nil != err {
			a.JSON(w, err)
		}

		a.JSON(w, data)

		break
	default:
		fmt.Fprintln(w, "Sorry, nothing here!")
	}
}
