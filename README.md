# guildedgo
A guilded.gg library in Go

## Getting started

```cmd
go get github.com/itschip/guildedgo
```

### Example

```go
func main() {
    
    // You might want to download or write your own functionality to get env variables
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
}
```
