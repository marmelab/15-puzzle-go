package events

import (
	"game"
	"github.com/nsf/termbox-go"
	"time"
	"unicode"
)

const SLEEP_DURATION time.Duration = 10

func detectGameCommand(event termbox.Event) byte {
	switch unicode.ToLower(event.Ch) {
	case 's':
		return game.ACTION_SHUFFLE
	case 'h':
		return game.ACTION_HELP
	case 0:
		switch event.Key {
		case termbox.KeyCtrlC, termbox.KeyEsc:
			return game.ACTION_QUIT
		case termbox.KeyArrowUp:
			return game.ACTION_MOVE_TOP
		case termbox.KeyArrowRight:
			return game.ACTION_MOVE_RIGHT
		case termbox.KeyArrowDown:
			return game.ACTION_MOVE_BOTTOM
		case termbox.KeyArrowLeft:
			return game.ACTION_MOVE_LEFT
		}
	}
	return 0
}

func KeyListener(inputChan chan byte) {
	for {
		event := termbox.PollEvent()

		switch event.Type {
		case termbox.EventKey:
			inputChan <- detectGameCommand(event)
		case termbox.EventInterrupt:
			inputChan <- game.ACTION_QUIT
			break
		case termbox.EventError:
			panic(event.Err)
		}
		time.Sleep(time.Millisecond * SLEEP_DURATION)
	}
}
