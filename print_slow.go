package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"

	"golang.org/x/term"
)

// Returns a random int within the range
// cribbed from https://gosamples.dev/random-numbers/
// No meaningful changes could really be made to this simple function which
// really ought to just be built-in
func randIntn(min, max int) int {
	return min + rand.Intn(max-min)
}

// Outputs a string rune-by-rune with a delay,
// stops delay if the user enters an escape character, currently
//
//	space
//	return
//	q
func printSlow(str string) {
	var wg sync.WaitGroup

	done_writing := make(chan bool, 1)
	hurry := make(chan bool, 1)

	wg.Add(1)
	go func() {
		done_printing := false
		defer wg.Done()
		for !done_printing {
			select {
			case done_printing = <-done_writing:
				return
			default:
				if readChar() {
					hurry <- true
					return
				}
			}
		}
	}()

	wg.Add(1)
	go func() {
		stop_delay := false
		defer wg.Done()
		for _, letter := range str {
			select {
			case stop_delay = <-hurry:
				fmt.Printf("%c", letter)
			default:
				if !stop_delay {
					time.Sleep(time.Duration(randIntn(25, 75)) * time.Millisecond)
				}
				fmt.Printf("%c", letter)
			}
		}
		done_writing <- true
	}()

	wg.Wait()
}

// Reads a character from stdin without printing it, then sends true if the
// char is an escape character
//
// taken from
// https://stackoverflow.com/questions/15159118/read-a-character-from-standard-input-in-go-without-pressing-enter
func readChar() bool {
	// switch stdin into 'raw' mode
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	// set read deadline for 50ms to match printChar timing
	read_deadline := time.Now().Add(time.Duration(50) * time.Millisecond)

	b := make([]byte, 1)
	os.Stdin.SetReadDeadline(read_deadline)
	_, err = os.Stdin.Read(b)
	if err != nil {
		fmt.Println(err)
	}

	escape_chars := "q\n\r "
	// fmt.Printf("\nthe char %q was hit\n", string(b[0]))
	return strings.Contains(escape_chars, string(b[0]))
}
