package events

import (
	"fmt"
	"renderer"
)

type Message struct {
	Content string
	Clear   bool
}

func RenderListener(msgChan chan Message) {
	for {
		msg := <-msgChan
		if msg.Clear {
			renderer.ClearTerminal()
		}
		fmt.Println(msg.Content)
	}
}
