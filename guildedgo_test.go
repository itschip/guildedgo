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

	_, err := c.Members.UpdateMemberNickname("", "")
	if err != nil {
		fmt.Println("Uh oh")
	}
}
