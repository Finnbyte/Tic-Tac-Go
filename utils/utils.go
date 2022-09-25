package utils

import (
	"bufio"
	"fmt"
	"os"
)

// Basic method of taking input, doesn't play well with space-separated lines.
func ScanlnWithPrompt(prompt string, cmdAddr *string) {
	fmt.Printf("%s", prompt)
	fmt.Scanln(cmdAddr)
}

// Used for space-separated input, so basically the whole game.
func BufioWithPrompt(prompt string, cmdAddr *string) {
	fmt.Printf("%s", prompt)
	var in *bufio.Reader = bufio.NewReader(os.Stdin)
	if line, _, err := in.ReadLine(); err != nil {
		panic(err)
	} else {
		*cmdAddr = string(line)
	}
}

func ClearOutput() {
	fmt.Printf("\033[2J\033[1;1H") // Special characters which corresponds to clearing screen.
	// This works on Linux and Windows, and thus is better than any hardcoded "cls" or "clear" command.
}
