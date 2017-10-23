package renderer

import (
	"os"
	"os/exec"	
)

func ClearTerminal() {
	clearCmd := exec.Command("clear")
	clearCmd.Stdout = os.Stdout
	clearCmd.Run()
}
