package guildedgo

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

func (c *Client) Open() {
	header := http.Header{}
	header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Token))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

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

	listening := make(chan struct{})

	go func() {
		defer close(listening)

		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println(err.Error())
				return
			}

			log.Printf("msg: %s", msg)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-listening:
			return

		case t := <-ticker.C:
			err := conn.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("write", err)
				return
			}
		case <-interrupt:
			log.Println("Interrupt")

			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close", err)
				return
			}

			select {
			case <-listening:
			case <-time.After(time.Second):
			}

			return
		}
	}
}
