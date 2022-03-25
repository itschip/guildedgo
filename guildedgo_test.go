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

	c := NewClient(config)

	message := &MessageObject{
		Content: "Hello Everyone!!",
	}

	msg, err := c.Channel.SendMessage("08dfae9c-6ecb-44b7-86ad-6812b495dd0c", message)
	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println(msg.Id)

	messages, err := c.Channel.GetMessages("08dfae9c-6ecb-44b7-86ad-6812b495dd0c", nil)
	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println((*messages)[0].Content)
}
