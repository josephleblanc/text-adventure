package main

//// Sources
//
// os.exit() usage
// https://gosamples.dev/exit/

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func prompt_enter() string {
	fmt.Printf("\n")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	defer check_quit(scanner)

	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	return scanner.Text()
}

func check_quit(scanner *bufio.Scanner) {
	if scanner.Text() == "quit" {
		os.Exit(0)
	}
}
