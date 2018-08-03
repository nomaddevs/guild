package beta

import (
	"net/http"

	"github.com/munsy/guild/errors"
	"github.com/munsy/guild/pkg/models"
)

// News creates a single news post or returns a set of posts, depending on the http method.
func (a *API) News(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		var nps models.NewsPosts

		err := nps.Read()

		if nil != err {
			e := &errors.Error{
				Message: err.Error(),
				Package: "api.beta",
				Type:    "API",
				Method:  "News",
			}
			a.Error(w, e)
			return
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
			e := &errors.Error{
				Message: err.Error(),
				Package: "api.beta",
				Type:    "API",
				Method:  "News",
			}
			a.Error(w, e)
			return
		}

		var nps models.NewsPosts

		err = nps.Read()

		if nil != err {
			e := &errors.Error{
				Message: err.Error(),
				Package: "api.beta",
				Type:    "API",
				Method:  "News",
			}
			a.Error(w, e)
			return
		}

		a.JSON(w, nps)
		break
	default:
		e := &errors.Error{
			Message: "default hit",
			Package: "api.beta",
			Type:    "API",
			Method:  "News",
		}
		a.Error(w, e)
		return
	}
}
