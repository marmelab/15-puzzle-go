package events

import (
	"os"
)

func QuitListener(interruptChan chan os.Signal, doneChan chan bool) {
	<-interruptChan
	doneChan <- false
}
