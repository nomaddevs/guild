package models

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type APIReader struct {
}

// Convert an HTTP response from a given URL to the supplied interface.
// This function expects the body to contain the associated JSON response
// from the given URL and will return an error if it fails to properly unmarshal.
func (reader *APIReader) Get(url string, v interface{}) error {
	if nil == v {
		return errors.New("No interface supplied.")
	}

	response, err := http.Get(url)
	if nil != err {
		return err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if nil != err {
		return err
	}

	err = json.Unmarshal([]byte(body), &v)
	if nil != err {
		return err
	}

	return nil
}
