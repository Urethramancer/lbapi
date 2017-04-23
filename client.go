package lbapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Client is the structure for reseller API access to LogicBoxes systems.
type Client struct {
	http.Client
	// URL of the API, usually https://httpapi.com.
	URL string
	// ID of the reseller.
	ID string
	// Key to authenticate with.
	Key string
}

// NewClient creates a client structure with a HTTP client for the specified API.
func NewClient(api string, resellerid int64, apikey string) *Client {
	if api == "" {
		api = APIURL
	}

	return &Client{
		Client: http.Client{Timeout: time.Second * 30},
		URL:    api,
		ID:     fmt.Sprintf("%d", resellerid),
		Key:    apikey,
	}
}

// GetReponse fetches a URL's JSON, decode it into a maplist
// and returns it as a map of strings.
func GetResponse(c http.Client, url string) (*maplist, error) {
	res, err := c.Get(url)
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

// PostResponse does pretty much the same as getResponse(),
// but with the POST method.
func PostResponse(c http.Client, url string) (*maplist, error) {
	res, err := c.Post(url, "", nil)
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
