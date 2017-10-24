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

func KeyListener(inputChan chan rune, doneChan chan bool) {
	for {
		event := termbox.PollEvent()

		switch event.Type {
		case termbox.EventKey:
			switch event.Key {
			case termbox.KeyCtrlC:
				doneChan <- false
				break
			case termbox.KeyEsc:
				doneChan <- false
				break
			case termbox.KeyArrowUp:
				inputChan <- 'Z'
				break
			case termbox.KeyArrowRight:
				inputChan <- 'D'
				break
			case termbox.KeyArrowDown:
				inputChan <- 'S'
				break
			case termbox.KeyArrowLeft:
				inputChan <- 'Q'
				break
			default:
				inputChan <- event.Ch
			}
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
		fmt.Println("Move the tiles with the arrow keys or press Esc to exit")
	}
}

func GameListener(doneChan chan bool, inputChan chan rune, gridChan chan game.Grid, startedGrid game.Grid) {
	grid := game.DeepCopyGrid(startedGrid)
	gridChan <- grid

	for {
		c := <-inputChan
		newCoords, err := game.CoordsFromDirection(grid, c)
		if err != nil {
			fmt.Println(err)
			continue
		}
		newGrid, err := game.Move(grid, newCoords)
		if err != nil {
			fmt.Println("The move is not possible, please try another direction")
			continue
		}
		grid = newGrid
		if reflect.DeepEqual(grid, startedGrid) {
			doneChan <- true
		}
		gridChan <- newGrid
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
	inputChan := make(chan rune)
	defer close(inputChan)
	gridChan := make(chan game.Grid)
	defer close(gridChan)

	clearTerminal()
	fmt.Print("Welcome to the 15 puzzle game ")
	time.Sleep(time.Second * 1)
	startedGrid := game.BuildGrid(4)

	signal.Notify(interruptChan, os.Interrupt, syscall.SIGTERM)
	go Quit(interruptChan, doneChan)
	go KeyListener(inputChan, doneChan)
	go RenderListener(gridChan)
	go GameListener(doneChan, inputChan, gridChan, startedGrid)

	success := <-doneChan
	if success {
		fmt.Println("\nGGWP, you solved the puzzle!")
	}

	fmt.Println("\nSee you soon :)")
	os.Exit(0)
}
