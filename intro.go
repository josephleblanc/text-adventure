package main

import (
	"fmt"
)

func print_title() {
	fmt.Printf("\n\n")
	center_text("Asleep at the Terminal")
	fmt.Printf("\n")
	intro_desc := "A Text Adventure Game"
	center_text(intro_desc)
	author := "by Brasides"
	center_text(author)
	fmt.Printf("\n")
}
