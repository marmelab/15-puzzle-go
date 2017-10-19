package renderer

import (
	"testing"
)

func TestHelloReturn(t *testing.T) {
	expectedMessage := "Welcome to the 15 puzzle game.\n"

	message := Hello()

	if message != expectedMessage {
		t.Error("The returned message is not the expected one")
	}
}
