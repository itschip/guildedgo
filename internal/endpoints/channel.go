package endpoints

import (
	"github.com/itschip/guildedgo/internal"
)

var (
	CreateMessageEndpoint = func(channelId string) string {
		return internal.GuildedApi + "/channels/" + channelId + "/messages"
	}
	GetChannelMessagesEndpoint = func(channelId string) string {
		return internal.GuildedApi + "/channels/" + channelId + "/messages"
	}
	GetChannelMessageEndpoint = func(channelId string, messageId string) string {
		return internal.GuildedApi + "/channels/" + channelId + "/messages/" + messageId
	}
)
