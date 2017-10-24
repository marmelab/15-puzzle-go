package main

import (
	e "events"
	"flag"
	"game"
	"github.com/nsf/termbox-go"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const SLEEP_DURATION time.Duration = 1
const DEFAULT_GRID_SIZE int = 4

func getGridSize() int {
	var gridSize int
	flag.IntVar(&gridSize, "size", DEFAULT_GRID_SIZE, "an int")
	flag.Parse()
	return gridSize
}

func main() {
	size := byte(getGridSize())
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	interruptChan := make(chan os.Signal, 2)
	defer close(interruptChan)
	doneChan := make(chan bool)
	defer close(doneChan)
	inputChan := make(chan byte)
	defer close(inputChan)
	msgChan := make(chan e.Message)
	defer close(msgChan)

	startedGrid := game.BuildGrid(size)

	signal.Notify(interruptChan, os.Interrupt, syscall.SIGTERM)
	go e.RenderListener(msgChan)
	msgChan <- e.Message{"Welcome to the 15 puzzle game", true}
	time.Sleep(time.Second * SLEEP_DURATION)

	go e.QuitListener(interruptChan, doneChan)
	go e.KeyListener(inputChan)
	go e.GameListener(doneChan, inputChan, msgChan, startedGrid)

	<-doneChan
	msgChan <- e.Message{"\nSee you soon :)", false}
	time.Sleep(time.Second * SLEEP_DURATION)
	os.Exit(0)
}
