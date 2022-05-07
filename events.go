package guildedgo

type EventHandler struct {
	events map[string]interface{}
}

type Event interface {
	Type() string

	Handle(*Client, interface{})
}

type messageCreatedEvent func(*Client, *ChatMessageCreated)

func (e messageCreatedEvent) Type() string {
	return "ChatMessageCreated"
}

func (e messageCreatedEvent) Handle(c *Client, i interface{}) {
	if t, ok := i.(*ChatMessageCreated); ok {
		e(c, t)
	}
}

func (c *Client) AddEventHander(h interface{}) {

}

func interfaceHandler(h interface{}) Event {
	switch sw := h.(type) {
	case func(*Client, *ChatMessageCreated):
		return messageCreatedEvent(sw)
	}

	return nil
}
