package main

import (
	"fmt"
	"game"
	"os"
	"os/exec"
	"renderer"
)

func clearTerminal() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func main() {
	clearTerminal()
	fmt.Println(renderer.Hello())
	grid := game.BuildGrid(4)
	fmt.Println(renderer.DrawGrid(grid))
	grid, err := game.MoveByValue(grid, byte(12))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(renderer.DrawGrid(grid))
	}
}
