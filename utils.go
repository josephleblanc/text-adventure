package main

//// Sources
//
// os.exit() usage
// https://gosamples.dev/exit/
//
// time.ticker usage
// https://gobyexample.com/tickers
// I've used their example code and made significant alterations to suite the
// program's purpse

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/term"
)

// Centers the input text in the terminal, padding with whitespace.
func center_text(text string) {
	width, _, err := term.GetSize(0)
	if err != nil {
		return
	}
	offset := (width + len(text)) / 2
	fmt.Printf("%*s\n", offset, text)
}

func clear() {
	_, height, err := term.GetSize(0)
	if err != nil {
		return
	}
	for range height {
		fmt.Printf("\n")
	}
}

// Prompts the user to continue description or exit program.
// Returns the string of user input.
func promptEnter() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	defer checkQuit(scanner)

	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	return scanner.Text()
}

// Checks if user has entered "quit" to exit program.
func checkQuit(scanner *bufio.Scanner) {
	if scanner.Text() == "quit" {
		os.Exit(0)
	}
}

// Writes an elipsis with delays between each `.`
func elipsis() {
	for i := 0; i < 3; i++ {
		fmt.Printf(".")
		time.Sleep(500 * time.Millisecond)
	}
}
