package guildedgo

import (
	"encoding/json"

	"github.com/itschip/guildedgo/internal/endpoints"
)

type ChannelService interface {
	SendMessage(channelId string, message *MessageObject) (*Message, error)
	GetMessages(channelId string, getObject *GetMessageObject) (*[]ChannelMessage, error)
}

type channelService struct {
	client *Client
}

var _ ChannelService = &channelService{}

func (cs *channelService) SendMessage(channelId string, message *MessageObject) (*Message, error) {
	endpoint := endpoints.CreateMessageEndpoint(channelId)

	resp, err := cs.client.PostRequest(endpoint, &message)
	if err != nil {
		return nil, err
	}

	var msg MessageResponse
	err = json.Unmarshal(resp, &msg)
	if err != nil {
		return nil, err
	}

	return &msg.Message, err
}

func (cs *channelService) GetMessages(channelId string, getObject *GetMessageObject) (*[]ChannelMessage, error) {
	endpoint := endpoints.GetChannelMessagesEndpoint(channelId)

	resp, err := cs.client.GetRequest(endpoint, nil)
	if err != nil {
		return nil, err
	}

	var msgs AllMessagesResponse
	err = json.Unmarshal(resp, &msgs)
	if err != nil {
		return nil, err
	}

	return &msgs.Messages, nil
}
