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
func centerText(text string) {
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
	defer checkHelp(scanner)

	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	return scanner.Text()
}

// prompt the user to continue with "c" or "continue"
// blocks the program until the user continues
func promptContinue() bool {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}
		if checkContinue(scanner) {
			return true
		} else {
			checkQuit(scanner)
			if checkHelp(scanner) {
				helpInfo()
			}
			centerText("Sorry, but I do not understand.")
			centerText("<Type \"continue\" or \"c\" to continue, or quit with \"quit\" or \"q\">")
		}
	}
}

// Checks if user has entered "quit" to exit program.
func checkQuit(scanner *bufio.Scanner) {
	keywords := [...]string{
		"quit",
		"q",
	}
	for _, str := range keywords {
		if str == scanner.Text() {
			os.Exit(0)
		}
	}
}

// Checks if user input contains a help request and prints the help message.
func checkHelp(scanner *bufio.Scanner) bool {
	keywords := [...]string{
		"help",
		"h",
	}
	for _, str := range keywords {
		if str == scanner.Text() {
			return true
		}
	}
	return false
}

func checkContinue(scanner *bufio.Scanner) bool {
	keywords := [...]string{
		"continue",
		"c",
	}
	for _, str := range keywords {
		if str == scanner.Text() {
			return true
		}
	}
	return false
}

// Writes an elipsis with delays between each `.`
func elipsis() {
	for i := 0; i < 3; i++ {
		fmt.Printf(".")
		time.Sleep(500 * time.Millisecond)
	}
}
