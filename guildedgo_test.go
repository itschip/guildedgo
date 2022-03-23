package guildedgo

import (
	"fmt"
	"log"
	"testing"

	"github.com/itschip/guildedgo/utils"
)

func TestNewClient(t *testing.T) {
	token := utils.GoDotEnvVariable("BOT_TOKEN")

	config := &Config{
		Token: token,
	}

	c := New(config)

	message := &MessageObject{
		Content: "Hello Everyone!!",
	}

	// wew, it works
	msg, err := c.SendChannelMessage("08dfae9c-6ecb-44b7-86ad-6812b495dd0c", message)
	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println(msg.Id)
}
