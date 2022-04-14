package guildedgo

import (
	"bytes"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)


func (c *Client) Open() {
	header := http.Header{}
	header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	
	conn, _, err := websocket.DefaultDialer.Dial("wss://api.guilded.gg/v1/websocket", header)
	if err != nil {
		log.Fatalln("Failed to connect to websocket: ", err.Error())
	}
	
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatalln(err.Error())
		}
	}()
	
	_, m, err := conn.ReadMessage()
	if err != nil {
		log.Fatalln("Failed to read message: ", err.Error())
	}
	
	m = bytes.TrimSpace(bytes.Replace(m, newline, space, -1))
	fmt.Println(string(m))
	
	listening := make(chan interface{})
	go c.beat(conn, listening, 22500)
}

func (c *Client) beat(conn *websocket.Conn, listening <-chan interface{}, interval time.Duration) {
	tick := time.NewTicker(interval * time.Millisecond)
	
	defer tick.Stop()
}