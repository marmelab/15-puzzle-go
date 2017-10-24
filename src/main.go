package main

import (
	e "events"
	"fmt"
	"game"
	"github.com/nsf/termbox-go"
	"os"
	"os/signal"
	"renderer"
	"syscall"
	"time"
)

func main() {
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
	gridChan := make(chan game.Grid)
	defer close(gridChan)

	renderer.ClearTerminal()
	fmt.Print("Welcome to the 15 puzzle game ")
	time.Sleep(time.Second * 1)
	startedGrid := game.BuildGrid(4)

	signal.Notify(interruptChan, os.Interrupt, syscall.SIGTERM)
	go e.QuitListener(interruptChan, doneChan)
	go e.KeyListener(inputChan)
	go e.RenderListener(gridChan)
	go e.GameListener(doneChan, inputChan, gridChan, startedGrid)

	success := <-doneChan
	if success {
		fmt.Println("\nGGWP, you solved the puzzle!")
	}

	fmt.Println("\nSee you soon :)")
	os.Exit(0)
}
