package main

import (
	"fmt"
	"game"
	"github.com/nsf/termbox-go"
	"os"
	"os/exec"
	"os/signal"
	"reflect"
	"renderer"
	"syscall"
	"time"
)

func clearTerminal() {
	clearCmd := exec.Command("clear")
	clearCmd.Stdout = os.Stdout
	clearCmd.Run()
}

func detectGameCommand(event termbox.Event) byte {
	switch event.Ch {
	case 'S', 's':
		return game.ACTION_SHUFFLE
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
	}
}

func Quit(interruptChan chan os.Signal, doneChan chan bool) {
	<-interruptChan
	doneChan <- false
}

func RenderListener(gridChan chan game.Grid) {
	for {
		grid := <-gridChan
		clearTerminal()
		fmt.Println(renderer.DrawGrid(grid))
		fmt.Println("Move the tiles with the arrow keys or or press (S) to shuffle or press Esc to exit")
	}
}

func GameListener(doneChan chan bool, inputChan chan byte, gridChan chan game.Grid, startedGrid game.Grid) {
	grid := game.DeepCopyGrid(startedGrid)
	gridChan <- grid
	for {
		action := <-inputChan
		if action == game.ACTION_QUIT {
			doneChan <- false
			return
		}
		
		if action == game.ACTION_SHUFFLE {
			clearTerminal()
			fmt.Println("Shuffling...")
			grid = game.Shuffle(grid)
		} else {
			newCoords, err := game.CoordsFromDirection(grid, action)
			if err != nil {
				fmt.Println(err)
				continue
			}
			grid, err = game.Move(grid, newCoords)
			if err != nil {
				fmt.Println("The move is not possible, please try another direction")
				continue
			}
			if reflect.DeepEqual(grid, startedGrid) {
				doneChan <- true
			}
		}
		gridChan <- game.DeepCopyGrid(grid)
	}
}

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

	clearTerminal()
	fmt.Print("Welcome to the 15 puzzle game ")
	time.Sleep(time.Second * 1)
	startedGrid := game.BuildGrid(4)

	signal.Notify(interruptChan, os.Interrupt, syscall.SIGTERM)
	go Quit(interruptChan, doneChan)
	go KeyListener(inputChan)
	go RenderListener(gridChan)
	go GameListener(doneChan, inputChan, gridChan, startedGrid)

	success := <-doneChan
	if success {
		fmt.Println("\nGGWP, you solved the puzzle!")
	}

	fmt.Println("\nSee you soon :)")
	os.Exit(0)
}
