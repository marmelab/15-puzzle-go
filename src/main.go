package main

import (
	"bufio"
	"fmt"
	"game"
	"os"
	"os/exec"
	"os/signal"
	"renderer"
	"syscall"
)

func clearTerminal() {
	clearCmd := exec.Command("clear")
	clearCmd.Stdout = os.Stdout
	clearCmd.Run()
}

func KeyListener(inputChan chan rune) {
	for {
		reader := bufio.NewReaderSize(os.Stdin, 1)
		c, _, _ := reader.ReadRune()
		inputChan <- c
	}
}

func RenderListener(gridChan chan game.Grid) {
	for {
		grid := <-gridChan
		clearTerminal()
		fmt.Println(renderer.DrawGrid(grid))
		fmt.Printf("Enter a direction (Z, Q, S or D) or press E to exit:\n> ")
	}
}

func GameListener(doneChan chan bool, inputChan chan rune, gridChan chan game.Grid, startedGrid game.Grid) {
	grid := game.DeepCopyGrid(startedGrid)
	gridChan <- grid

	for {
		c := <-inputChan
		if c == 'e' || c == 'E' {
			doneChan <- true
		} else {
			newCoords, _ := game.CoordsFromDirection(grid, c)
			newGrid, err := game.Move(grid, newCoords)

			if err == nil {
				grid = newGrid
				gridChan <- newGrid
			} else {
				fmt.Printf("Wrong command, please retry\n> ")
			}
		}
	}
}

func Quit(interruptChan chan os.Signal, doneChan chan bool) {
	<-interruptChan
	doneChan <- true
}

func main() {

	interruptChan := make(chan os.Signal, 2)
	doneChan := make(chan bool)
	inputChan := make(chan rune)
	gridChan := make(chan game.Grid)

	clearTerminal()
	fmt.Println(renderer.Hello())

	startedGrid := game.BuildGrid(4)

	signal.Notify(interruptChan, os.Interrupt, syscall.SIGTERM)
	go Quit(interruptChan, doneChan)
	go KeyListener(inputChan)
	go RenderListener(gridChan)
	go GameListener(doneChan, inputChan, gridChan, startedGrid)

	<-doneChan
	fmt.Println("\nGoodbye")
	os.Exit(0)
}
