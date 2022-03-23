package guildedgo

type Client struct {
	Token string
}

type Config struct {
	Token string
}

func New(config *Config) *Client {
	c := &Client{Token: config.Token}

	return c
}
