package beta

import (
	"fmt"
	"net/http"
)

func handleMakeNewsPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Sorry, nothing here!")
}
