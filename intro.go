package main

import (
	"fmt"
	"text-adventure/utils"
)

// Print game title, short description, and author in the center of the
// terminal.
func printTitle() {
	fmt.Printf("\n\n")
	utils.CenterText("Asleep at the Terminal")
	fmt.Printf("\n")
	intro_desc := "A Text Adventure Game"
	utils.CenterText(intro_desc)
	author := "by Brasides"
	utils.CenterText(author)
	fmt.Printf("\n")
}
