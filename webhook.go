package guildedgo

type Webhook struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ServerID  string `json:"serverId"`
	ChannelID string `json:"channelId"`
	CreatedAt string `json:"createdAt"`
	CreatedBy string `json:"createdBy"`
	DeletedAt string `json:"deletedAt"`
	Token     string `json:"string"`
}

type ChatMessageCreated struct {
	OP int                    `json:"op"`
	T  string                 `json:"t"`
	D  ChatMessageCreatedData `json:"d"`
}

type ChatMessageCreatedData struct {
	ServerID string      `json:"serverId"`
	Message  ChatMessage `json:"message"`
}
