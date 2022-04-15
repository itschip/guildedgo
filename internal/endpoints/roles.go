package endpoints

import "github.com/itschip/guildedgo/internal"

var (
	GroupMemberEndpoint = func(groupId string, userId string) string {
		return internal.GuildedApi + "/groups/" + groupId + "/members/" + userId
	}
)