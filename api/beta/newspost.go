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
	db := &MariaDB{
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
			fmt.Println("[SERVER][GET] - Error from api.beta.News()")
			a.JSON(w, err)
		}

		fmt.Println("[SERVER][GET] - Succeeded from api.beta.News()")

		a.JSON(w, data)

		break
	case "POST":
		title := r.FormValue("title")
		body := r.FormValue("body")
		author := r.FormValue("author")

		err := db.WriteNewsPost(title, body, author)

		if nil != err {
			fmt.Println("[SERVER][POST] - Error from api.beta.News()")
			a.JSON(w, err)
		}

		fmt.Println("[SERVER][POST] - Succeeded from api.beta.News()")

		data, err := db.ReadNewsPosts()

		if nil != err {
			fmt.Println("[SERVER][GET] - Error from api.beta.News()")
			a.JSON(w, err)
		}

		a.JSON(w, data)

		break
	default:
		fmt.Fprintln(w, "Sorry, nothing here!")
	}
}
