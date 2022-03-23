package guildedgo

import (
	"encoding/json"
	"fmt"

	"github.com/itschip/guildedgo/internal/endpoints"
)

type Channel struct {
}

func (c *Client) SendChannelMessage(channelId string, message *MessageObject) (*Message, error) {
	endpoint := endpoints.CreateMessageEndpoint(channelId)

	resp, err := c.PostRequest(endpoint, &message)
	if err != nil {
		return nil, err
	}

	var msg MessageReponse
	err = json.Unmarshal(resp, &msg)
	if err != nil {
		return nil, err
	}

	fmt.Println(&msg)

	return &msg.Message, err
}
