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
	fmt.Println(renderer.BuildGrid(grid))
}
