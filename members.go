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
	UpdateMemberNickname(userId string, nickname string) (*NicknameResponse, error)
	GetServerMember(serverId string, userId string) (*ServerMember, error)
	KickMember(userId string) error
}

type membersService struct {
	client *Client
}

var _ MembersService = &membersService{}

func (ms *membersService) UpdateMemberNickname(userId string, nickname string) (*NicknameResponse, error) {
	endpoint := endpoints.UpdateMemberNicknameEndpoint(ms.client.ServerID, userId)

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

func (ms *membersService) KickMember(userId string) error {
	endpoint := endpoints.ServerMemberEndpoint(ms.client.ServerID, userId)

	// No response. Maybe ther'e be one in the future
	_, err := ms.client.DeleteRequest(endpoint)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}

	return nil
}

func (ms *membersService) BanMember(userId string, reason string) (*ServerMemberBan, error) {
	endpoint := endpoints.MemberBanEndpoint(ms.client.ServerID, userId)

	// No need to build a struct here
	body := map[string]string{
		"reason": reason,
	}

	resp, err := ms.client.PostRequest(endpoint, body)
	if err != nil {
		log.Fatalln(err.Error())
		return nil, err
	}

	var ban ServerMemberBan
	err = json.Unmarshal(resp, &ban)
	if err != nil {
		log.Fatalln(err.Error())
		return nil, err
	}

	return &ban, nil
}

func (ms *membersService) IsMemberBanned(userId string) (*ServerMemberBan, error) {
	// Do we want to use the serverID from the config, or manually input it?
	endpoint := endpoints.ServerMemberEndpoint(ms.client.ServerID, userId)

	resp, err := ms.client.GetRequest(endpoint)
	if err != nil {
		log.Fatalln(err.Error())
		return nil, err
	}

	var ban ServerMemberBan
	err = json.Unmarshal(resp, &ban)
	if err != nil {
		log.Fatalln(err.Error())
		return nil, err
	}

	return &ban, nil
}

func (ms *membersService) UnbanMember(userId string) error {
	endpoint := endpoints.MemberBanEndpoint(ms.client.ServerID, userId)

	_, err := ms.client.DeleteRequest(endpoint)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}

	return nil
}

func (ms *membersService) GetServerMembers() (*[]ServerMemberSummary, error) {
	endpoint := endpoints.MemberEndpoint(ms.client.ServerID)

	resp, err := ms.client.GetRequest(endpoint)
	if err != nil {
		log.Fatalln(err.Error())
		return nil, err
	}

	var members []ServerMemberSummary
	err = json.Unmarshal(resp, &members)
	if err != nil {
		log.Fatalln(err.Error())
		return nil, err
	}

	return &members, nil
}
