package guildedgo

import (
	"fmt"
	"log"
	"testing"

	"github.com/itschip/guildedgo/utils"
)

func TestNewClient(t *testing.T) {
	token := utils.GoDotEnvVariable("BOT_TOKEN")
	channelId := utils.GoDotEnvVariable("TEST_CHANNEL_ID")

	config := &Config{
		Token: token,
	}

	c := NewClient(config)

	message := &MessageObject{
		Content: "Hello Everyone!!",
	}

	msg, err := c.Channel.SendMessage(channelId, message)
	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println(msg.Id)

	messages, err := c.Channel.GetMessages(channelId, nil)
	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println((*messages)[0].Content)
}
