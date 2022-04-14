package guildedgo

import (
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
	
	c.Open()
}
