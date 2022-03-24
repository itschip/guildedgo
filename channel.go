package guildedgo

import (
	"encoding/json"
	"fmt"

	"github.com/itschip/guildedgo/internal/endpoints"
)

type ChannelService interface {
	SendChannelMessage(channelId string, message *MessageObject) (*Message, error)
}

type channelService struct {
	client *Client
}

var _ ChannelService = &channelService{}

func (cs *channelService) SendChannelMessage(channelId string, message *MessageObject) (*Message, error) {
	endpoint := endpoints.CreateMessageEndpoint(channelId)

	resp, err := cs.client.PostRequest(endpoint, &message)
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
