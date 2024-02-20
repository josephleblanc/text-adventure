package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
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
// stops delay if the user enters an escape character.
// Currently works, but is commented out while I try another method below.
//
//	space
//	return
//	q
// func printSlow(str string) {
// 	var wg sync.WaitGroup
//
// 	done_writing := make(chan bool, 1)
// 	hurry := make(chan bool, 1)
//
// 	wg.Add(1)
// 	go func() {
// 		done_printing := false
// 		defer wg.Done()
// 		for !done_printing {
// 			select {
// 			case done_printing = <-done_writing:
// 				return
// 			default:
// 				if readChar() {
// 					hurry <- true
// 					return
// 				}
// 			}
// 		}
// 	}()
//
// 	wg.Add(1)
// 	go func() {
// 		stop_delay := false
// 		defer wg.Done()
// 		for _, letter := range str {
// 			select {
// 			case stop_delay = <-hurry:
// 				fmt.Printf("%c", letter)
// 			default:
// 				if !stop_delay {
// 					time.Sleep(time.Duration(randIntn(25, 75)) * time.Millisecond)
// 				}
// 				fmt.Printf("%c", letter)
// 			}
// 		}
// 		done_writing <- true
// 	}()
//
// 	wg.Wait()
// }

func printSlow(str string) {
	for i, c := range str {
		os.Stdin.SetReadDeadline(time.Now().Add(time.Duration(randIntn(25, 75)) * time.Millisecond))
		hurry := readChar()
		if hurry {
			fmt.Printf("%s", str[i:])
			return
		}
		fmt.Printf("%c", c)
	}
}

//// Currently does not work, I am trying to make `SyscallConn` work to allow
//`SetReadDeadline` a valid file/pipe
//
// Reads a character from stdin without printing it, then sends true if the
// char is an escape character
//
// taken from
// https://stackoverflow.com/questions/15159118/read-a-character-from-standard-input-in-go-without-pressing-enter

func readChar() bool {
	// switch stdin into 'raw' mode
	sysconn_fd, err := os.Stdin.SyscallConn()
	if err != nil {
		fmt.Println(err)
	}

	b := make([]byte, 1)
	var oldState *term.State
	err = sysconn_fd.Control(func(fd uintptr) {
		oldState, err = term.MakeRaw(int(fd))
		if err != nil {
			fmt.Println(err)
		}
	})
	if err != nil {
		fmt.Println(err)
	}

	_, err = os.Stdin.Read(b)
	if err != nil {
		fmt.Println(err)
	}

	err = sysconn_fd.Control(func(fd uintptr) {
		term.Restore(int(fd), oldState)
	})
	if err != nil {
		fmt.Println(err)
	}

	escape_chars := "q\n\r "
	return strings.Contains(escape_chars, string(b[0]))
}
