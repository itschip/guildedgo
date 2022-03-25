package guildedgo

import (
	"encoding/json"

	"github.com/itschip/guildedgo/internal"
)

func (c *Client) PostRequest(endpoint string, body interface{}) ([]byte, error) {
	jsonBody, _ := json.Marshal(&body)

	resp, err := internal.DoRequest("POST", endpoint, jsonBody, c.Token)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetRequest(endpoint string, body interface{}) ([]byte, error) {
	resp, err := internal.DoRequest("GET", endpoint, nil, c.Token)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
