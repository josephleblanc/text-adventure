package main

import (
	"fmt"

	"golang.org/x/term"
)

func print_title() {
	intro_msg := "Asleep at the Terminal"
	center_text(intro_msg)
	fmt.Printf("\n")
	intro_desc := "A Text Adventure Game"
	center_text(intro_desc)
	author := "by Brasides"
	center_text(author)
	fmt.Printf("\n")
}

func center_text(text string) {
	width, _, err := term.GetSize(0)
	if err != nil {
		return
	}
	offset := (width + len(text)) / 2
	fmt.Printf("%*s\n", offset, text)
}
