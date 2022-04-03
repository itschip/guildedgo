package endpoints

import "github.com/itschip/guildedgo/internal"

var (
	MemberEndpoint = func(serverId string) string {
		return internal.GuildedApi + "/servers/" + serverId + "/members"
	}
	ServerMemberEndpoint = func(serverId string, userId string) string {
		return internal.GuildedApi + "/servers/" + serverId + "/members/" + userId
	}
	UpdateMemberNicknameEndpoint = func(serverId string, userId string) string {
		return ServerMemberEndpoint(serverId, userId) + "/nickname"
	}
)