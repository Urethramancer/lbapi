package lbapi

import (
	"encoding/json"
	"net/http"
	"time"
)

type Client struct {
	http.Client
	URL string
	ID  string
	Key string
}

// NewClient creates a client structure with a HTTP client for the specified API.
func NewClient(api, resellerid, apikey string) *Client {
	if api == "" {
		api = APIURL
	}

	return &Client{
		Client: http.Client{Timeout: time.Second * 10},
		URL:    api,
		ID:     resellerid,
		Key:    apikey,
	}
}

// getReponse is the internal call to fetch a URL's JSON,
// decode it into a maplist and hand it over to massage into
// a proper structure.
func (c *Client) getResponse(url string) (*maplist, error) {
	res, err := c.Client.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	var list maplist
	err = decoder.Decode(&list)
	if err != nil {
		return nil, err
	}

	return &list, nil
}
