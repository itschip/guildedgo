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
	fmt.Println(msg.Id, msg.ChannelId)
	
	newMessage := &MessageObject{
		Content: "Bye Everyone!!",
	}
	
	newMsg,_ := c.Channel.UpdateChannelMessage(msg.ChannelId, msg.Id, newMessage)
	
	fmt.Println(newMsg.Id)
	
	err = c.Channel.DeleteChannelMessage(newMsg.ChannelId, newMsg.Id)
	if err != nil {
		log.Println(err.Error())
	}
}
