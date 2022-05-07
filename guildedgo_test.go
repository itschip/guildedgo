package guildedgo

import (
	"fmt"
	"testing"

	"github.com/itschip/guildedgo/utils"
)

func TestNewClient(t *testing.T) {
	token := utils.GoDotEnvVariable("BOT_TOKEN")

	config := &Config{
		ServerID: "",
		Token:    token,
	}

	c := NewClient(config)

	c.AddEventHander(messageCreate)

	c.Open()
}

func messageCreate(c *Client, m *ChatMessageCreated) {
	fmt.Println("Helllo message")
}
