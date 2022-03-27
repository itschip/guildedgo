package guildedgo

import (
	"encoding/json"
	
	"github.com/itschip/guildedgo/internal/endpoints"
)

type ChannelService interface {
	SendMessage(channelId string, message *MessageObject) (*ChatMessage, error)
	GetMessages(channelId string, getObject *GetMessagesObject) (*[]ChannelMessage, error)
	GetMessage(channelId string, messageId string) (*ChatMessage, error)
}

type channelService struct {
	client *Client
}

var _ ChannelService = &channelService{}

func (cs *channelService) SendMessage(channelId string, message *MessageObject) (*ChatMessage, error) {
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

// GetMessages TODO: add support for params
func (cs *channelService) GetMessages(channelId string, getObject *GetMessagesObject) (*[]ChannelMessage, error) {
	endpoint := endpoints.GetChannelMessagesEndpoint(channelId)
	
	resp, err := cs.client.GetRequest(endpoint)
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

// GetMessage Get a message from a channel
func (cs *channelService) GetMessage(channelId string, messageId string) (*ChatMessage, error) {
	endpoint := endpoints.GetChannelMessageEndpoint(channelId, messageId)
	
	resp, err := cs.client.GetRequest(endpoint)
	if err != nil {
		return nil, err
	}
	
	var msg MessageResponse
	err = json.Unmarshal(resp, &msg)
	if err != nil {
		return nil, err
	}
	
	return &msg.Message, nil
}
