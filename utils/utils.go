package utils

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
	"strings"
	"text-adventure/myprint"
	"text-adventure/mytypes"
	"time"

	"golang.org/x/term"
)

// Centers the input text in the terminal, padding with whitespace.
func CenterText(text string) {
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
func PromptEnter() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	defer CheckQuit(scanner)
	defer CheckHelp(scanner)

	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	return scanner.Text()
}

// prompt the user to continue with "c" or "continue"
// blocks the program until the user continues
func PromptContinue() bool {
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
			CheckQuit(scanner)
			if CheckHelp(scanner) {
				HelpInfo()
			}
			CenterText("Sorry, but I do not understand.")
			CenterText("<Type \"continue\" or \"c\" to continue, or quit with \"quit\" or \"q\">")
		}
	}
}

// prompt the user to continue with "c" or "continue"
// blocks the program until the user continues
func PromptName(user_input *mytypes.UserInput) mytypes.PlayerData {
	fmt.Printf("\nName: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	user_input.Selection = scanner.Text()
	// name := scanner.Text()
	verifyUserInput(scanner, user_input)
	return mytypes.NewPlayerData(user_input.Selection)
}

func PromptNav(p *mytypes.Player, m *mytypes.Map) bool {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}
		if tryLook(scanner, p, m) {
			return true
		} else if tryTravel(scanner, p, m) {
			return true
		} else if tryCurrentRoom(scanner, p, m) {
			return true
		} else {
			CheckQuit(scanner)
			if CheckHelp(scanner) {
				CenterText("<Type \"look\" or \"l\" and a direction (\"north\", \"east\", \"south\", \"west\"), or \"go\" and a direction, or type \"h\" or \"help\" for help or \"q\" or \"quit\" to quit.>")
			}
			CenterText("Sorry, but I do not understand.")
			HelpInfo()
			return false
		}
	}
}

func tryCurrentRoom(scanner *bufio.Scanner, p *mytypes.Player, m *mytypes.Map) bool {
	keywords := [...]string{
		"room",
		"r",
	}

	input := strings.Trim(scanner.Text(), " ")
	split := strings.Split(input, " ")

	for _, str := range keywords {
		if len(split) == 1 && split[0] == str {
			m.CurrentRoom(p)
			return true
		}
	}
	return false
}

func tryLook(scanner *bufio.Scanner, p *mytypes.Player, m *mytypes.Map) bool {
	keywords := [...]string{
		"look",
		"l",
	}

	input := strings.Trim(scanner.Text(), " ")
	split := strings.Split(input, " ")

	for _, str := range keywords {
		if len(split) == 2 && split[0] == str {
			m.Look(p, split[1])
			return true
		}
	}
	return false
}

func tryTravel(scanner *bufio.Scanner, p *mytypes.Player, m *mytypes.Map) bool {
	keywords := [...]string{
		"go",
		"g",
	}

	input := strings.Trim(scanner.Text(), " ")
	split := strings.Split(input, " ")

	for _, str := range keywords {
		if len(split) == 2 && split[0] == str {
			m.Travel(p, split[1])
			return true
		}
	}
	return false
}

func verifyUserInput(scanner *bufio.Scanner, user_input *mytypes.UserInput,
) string {
	myprint.PrintSlow(user_input.PreVerify + user_input.Selection + user_input.PostVerify)
	fmt.Println()
	CenterText("<Type \"y\" to confirm, or \"n\" to enter again.>")
	for {
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}
		CheckQuit(scanner)
		if CheckHelp(scanner) {
			HelpInfo()
		}
		switch scanner.Text() {
		case "y":
			return user_input.Selection
		case "n":
			fmt.Println()
			CenterText("<Please enter your name>")
			fmt.Printf("%s: ", user_input.FieldName)
			scanner.Scan()
			user_input.Selection = scanner.Text()
			return verifyUserInput(scanner, user_input)
		}

	}
}

// Display a general purpose help message to user when they have entered an invalid input.
func HelpInfo() {
	CenterText("<For help, type \"help\" or \"h\". To exit the program, type \"quit\" or \"q\".>")
}

// Checks if user has entered "quit" to exit program.
func CheckQuit(scanner *bufio.Scanner) {
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
func CheckHelp(scanner *bufio.Scanner) bool {
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
func Ellipsis() {
	for i := 0; i < 3; i++ {
		fmt.Printf(".")
		time.Sleep(500 * time.Millisecond)
	}
}

// Remove empty strings in array by copying any non-empty strings to new array.
func RemoveEmpty(s_array []string) []string {
	new_array := make([]string, 0)
	for _, s := range s_array {
		if s != "" {
			new_array = append(new_array, s)
		}
	}
	return new_array
}
