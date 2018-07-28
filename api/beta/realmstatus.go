package beta

// RealmStatus handles realm status data
func (a *API) RealmStatus(w http.ResponseWriter, r *http.Request) {
	client, err := battlenet.WoWClient(a.settings.BlizzardSettings(), a.key)

	if nil != err {
		fmt.Fprintln(w, "There was an error :(")
		fmt.Println(w, err.Error())
		return
	}

	switch r.Method {
	case "GET":
		response, err := client.RealmStatus()

		if nil != err {
			fmt.Fprintln(w, "There was an error :(")
			fmt.Println(w, err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response.Data)
		break
	default:
		fmt.Fprintln(w, "Sorry, nothing here!")
	}
}
