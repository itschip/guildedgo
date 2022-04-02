package guildedgo

import (
	"encoding/json"
	"fmt"

	"github.com/itschip/guildedgo/internal/endpoints"
)

type ChannelService interface {
	SendMessage(channelId string, message *MessageObject) (*ChatMessage, error)
	GetMessages(channelId string, getObject *GetMessagesObject) (*[]ChannelMessage, error)
	GetMessage(channelId string, messageId string) (*ChatMessage, error)
	UpdateChannelMessage(channelId string, messageId string, newMessage *MessageObject) (*ChatMessage, error)
	DeleteChannelMessage(channelId string, messageId string) error
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

func (cs *channelService) UpdateChannelMessage(channelId string, messageId string, newMessage *MessageObject) (*ChatMessage, error) {
	endpoint := endpoints.UpdateChannelMessageEndpoint(channelId, messageId)

	resp, err := cs.client.PutRequest(endpoint, &newMessage)
	if err != nil {
		return nil, err
	}

	var msg MessageResponse
	err = json.Unmarshal(resp, &msg)
	if err != nil {
		return nil, err
	}

	fmt.Println(msg)
	return &msg.Message, err
}

// GetMessages TODO: add support for params
func (cs *channelService) GetMessages(channelId string, getObject *GetMessagesObject) (*[]ChannelMessage, error) {
	endpoint := endpoints.GetChannelMessagesEndpoint(channelId)

	resp, err := cs.client.GetRequest(endpoint)
	if err != nil {
		return nil, err
	}

	// Abstract this functionality in GetRequest, as for the rest below and above
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

func (cs *channelService) DeleteChannelMessage(channelId string, messageId string) error {
	endpoint := endpoints.GetChannelMessageEndpoint(channelId, messageId)

	_, err := cs.client.DeleteRequest(endpoint)
	if err != nil {
		return err
	}

	return nil
}
