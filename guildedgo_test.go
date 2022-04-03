package guildedgo

import (
	"testing"

	"github.com/itschip/guildedgo/utils"
)

func TestNewClient(t *testing.T) {
	token := utils.GoDotEnvVariable("BOT_TOKEN")

	config := &Config{
		ServerID: "gRG4yLYl",
		Token:    token,
	}

	c := NewClient(config)

	_ = c.Members.KickMember("dKR1a6e4")
}
