# guildedgo
A guilded.gg library in Go

^ That's it for now

```go
func main() {
	token := utils.GoDotEnvVariable("BOT_TOKEN")
	c := New(&Config{
		Token: token,
	})

	message := &MessageObject{
		Content: "Hello Everyone!!",
	}
	msg, err := c.SendChannelMessage("08dfae9c-6ecb-44b7-86ad-6812b495dd0c", message)
	if err != nil {
		log.Println(err.Error())
	}
}
```
