package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	apiKey   string
	apiToken string
}

const (
	BASE_URL = "https://api.trello.com/1/"
)

// NewClient creates a new Trello API client
func NewClient(apiKey, apiToken string) Client {
	return Client{apiKey, apiToken}
}

// FetchListCards fetches all cards from the given list
func (c Client) FetchListCards(listId string) (ListApiResponse, error) {
	url := fmt.Sprintf("%slist/%s/cards?key=%s&token=%s", BASE_URL, listId, c.apiKey, c.apiToken)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve cards: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response body: %w", err)
	}

	res := new(ListApiResponse)
	err = json.Unmarshal(body, &res)
	return *res, err
}
