package guildedgo

import (
	"encoding/json"
	"log"

	"github.com/itschip/guildedgo/internal/endpoints"
)

type ServerMember struct {
	User     User   `json:"user"`
	RoleIds  []int  `json:"roleIds"`
	Nickname string `json:"nickname"`
	JoinedAt string `json:"joinedAt"`
}

type ServerMemberSummary struct {
	User    User  `json:"user"`
	RoleIds []int `json:"roleIds"`
}

type User struct {
	Id        string `json:"id"`
	Type      string `json:"type"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
}

type UserSummary struct {
	Id   string `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
}

type ServerMemberBan struct {
	User      User   `json:"user"`
	Reason    string `json:"reason"`
	CreatedAt string `json:"createdAt"`
	/*
		The ID of the user who banned
	*/
	CreatedBy string `json:"created_by"`
}

type NicknameResponse struct {
	Nickname string `json:"nickname"`
}

type ServerMemberResponse struct {
	Member ServerMember `json:"member"`
}

type MembersService interface {
	UpdateMemberNickname(serverId string, userId string, nickname string) (*NicknameResponse, error)
	GetServerMember(serverId string, userId string) (*ServerMember, error)
}

type membersService struct {
	client *Client
}

var _ MembersService = &membersService{}

// TODO: Fix some forbidden error
func (ms *membersService) UpdateMemberNickname(serverId string, userId string, nickname string) (*NicknameResponse, error) {
	endpoint := endpoints.UpdateMemberNicknameEndpoint(serverId, userId)

	body := &NicknameResponse{
		Nickname: nickname,
	}

	resp, err := ms.client.PutRequest(endpoint, body)

	if err != nil {
		log.Fatalln(err.Error())
		return nil, err
	}

	var nick NicknameResponse

	err = json.Unmarshal(resp, &nick)
	if err != nil {
		log.Fatalln(err.Error())
		return nil, err
	}

	return &nick, nil
}

func (ms *membersService) GetServerMember(serverId string, userId string) (*ServerMember, error) {
	endpoint := endpoints.ServerMemberEndpoint(serverId, userId)

	resp, err := ms.client.GetRequest(endpoint)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var member ServerMemberResponse
	err = json.Unmarshal(resp, &member)
	if err != nil {
		log.Fatalln(err.Error())
		return nil, err
	}

	return &member.Member, nil
}
