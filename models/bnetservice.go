package models

import(
	"errors"
	"io/ioutil"
	"encoding/json"
	"net/http"

	bnet "github.com/mitchellh/go-bnet"
)

// BnetService is a wrapper for an AccountService.
// See: https://github.com/mitchellh/go-bnet/blob/master/account.go
type BnetService struct {
	client *BnetClient
}

// BnetClient is a wrapper for a go-bnet client.
// See: https://github.com/mitchellh/go-bnet/blob/master/client.go
type BnetClient struct {
	client *bnet.Client
}

// NewBnetClient builds a new BnetClient wrapper.
func NewBnetClient(region string, c *http.Client) *BnetClient {
	return &BnetClient{
		client: bnet.NewClient(region, c),
	}
}

// NewRequest creates a new Battle.net API request as part of the original go-bnet client wrapper.
// See: https://github.com/mitchellh/go-bnet/blob/master/client.go
func (bnc *BnetClient) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	return bnc.client.NewRequest(method, urlStr, body)
}

// Do sends a Battle.net API request and returns the API response as part of the original go-bnet client wrapper.
func (bnc *BnetClient) Do(req *http.Request, v interface{}) (*bnet.Response, error) {
	return bnc.client.Do(req, v)
}

// Account returns a BnetService pointer that points to the client.
func (bnc *BnetClient) Account() *BnetService {
	return &BnetService{
		client: bnc,
	}
}

// Convert an HTTP response from a given URL to the supplied interface.
// This function expects the body to contain the associated JSON response
// from the given URL and will return an error if it fails to properly unmarshal.
func Get(url string, v interface{}) error {
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
