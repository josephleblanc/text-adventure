package main

import (
	"fmt"
)

// Print game title, short description, and author in the center of the
// terminal.
func printTitle() {
	fmt.Printf("\n\n")
	centerText("Asleep at the Terminal")
	fmt.Printf("\n")
	intro_desc := "A Text Adventure Game"
	centerText(intro_desc)
	author := "by Brasides"
	centerText(author)
	fmt.Printf("\n")
}
