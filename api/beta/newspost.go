package beta

import (
	"fmt"
	"net/http"

	//"github.com/munsy/guild/config"
	//"github.com/munsy/guild/database"
	"github.com/munsy/guild/pkg/models"
)

// News creates a single news post or returns a set of posts, depending on the http method.
func (a *API) News(w http.ResponseWriter, r *http.Request) {
	/*
		db := &database.MariaDB{
			Username:       config.DBUsername,
			Unixsocketpath: config.DBUnixsocketpath,
			Password:       config.DBPassword,
			Host:           config.DBHost,
			Port:           config.DBPort,
			Database:       config.DBName,
			Charset:        config.DBCharset,
		}
	*/

	switch r.Method {
	case "GET":
		//data, err := db.ReadNewsPosts()

		var nps models.NewsPosts

		err := nps.Read()

		if nil != err {
			a.JSON(w, err)
		}

		a.JSON(w, nps)

		break
	case "POST":
		title := r.FormValue("post_title")
		body := r.FormValue("post_body")
		author := r.FormValue("post_author")

		np := &models.NewsPost{
			Title:  title,
			Body:   body,
			Author: author,
		}

		err := np.Save()

		if nil != err {
			a.JSON(w, err)
		}

		var nps models.NewsPosts

		err = nps.Read()

		if nil != err {
			a.JSON(w, err)
		}

		a.JSON(w, nps)

		break
	default:
		fmt.Fprintln(w, "Sorry, nothing here!")
	}
}
