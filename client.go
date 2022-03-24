package guildedgo

import "net/http"

type Client struct {
	Token  string
	client *http.Client

	Channel ChannelService
}

type Config struct {
	Token string
}

func NewClient(config *Config) *Client {
	c := &Client{
		Token:  config.Token,
		client: http.DefaultClient,
	}

	c.Channel = &channelService{client: c}

	return c
}
