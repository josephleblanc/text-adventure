package puzzles

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"text-adventure/mytypes"
	"text-adventure/utils"
)

var modus_ponens_example string = `(A)		 True:  I am a man/woman
(B)		 ?:     I am mortal
(C: A->B)	 True:  If I am a man/woman, then I am mortal.
mp a c
(A)		 True:  I am a man/woman
(C: A->B)	 True:  If I am a man/woman, then I am mortal.
 .
. .
(B)		 True:  I am mortal`

// A complete list of operations with a short description of each
var ops_list [][]string = [][]string{
	{"mp", "modus ponens: Given statement A true and the implication A->B true, therefore B true"},
}

func PromptTool(puz *Puzzle, backup *Puzzle, player *mytypes.Player) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	// Check for user commands not directly related to the puzzle
	utils.CheckQuit(scanner)
	utils.CheckHelp(scanner)
	CheckOps(scanner, player)
	CheckDetails(scanner)
	CheckReset(puz, backup, scanner)
	CheckStatus(puz, scanner)

	input := strings.Trim(scanner.Text(), " ")
	split := strings.Split(input, " ")

	if len(split) == 4 && (split[0] == "modus" && split[1] == "ponens") {
		// check for modus ponens command (long version)
		vals := HandleModusPonens(puz, split[2], split[3])
		for _, v := range vals {
			fmt.Println(v)
		}
	} else if len(split) == 3 && split[0] == "mp" {
		// check for modus ponens command (short version)
		vals := HandleModusPonens(puz, split[1], split[2])
		for _, v := range vals {
			fmt.Println(v)
		}
	} else if len(split) == 1 && split[0] == "" {
		PuzzleHelp()
	}
}

func PuzzleHelp() {
	utils.CenterText("<You are currently in a puzzle.>")
	fmt.Println("Try applying the tools you have to solve the puzzle, or you may type \"operations\" or \"ops\" to see available operations, or \"definitions\" or \"defs\" to see collected definitions.")
	fmt.Println("To see the puzzle status, enter \"status\". To reset the puzzle, type \"reset\"")
}

func CheckReset(puzzle *Puzzle, backup *Puzzle, scanner *bufio.Scanner) {
	if scanner.Text() == "reset" {
		puzzle = backup
	}
}

func CheckStatus(puzzle *Puzzle, scanner *bufio.Scanner) {
	if scanner.Text() == "puzzle" {
		puzzle.Status()
	}
}

// Checks the buffer for keywords related to user commands requesting a list of available logical operations.
// Which operations are available is determined by the user's progress so far.
func CheckOps(scanner *bufio.Scanner, player *mytypes.Player) {
	keywords := [...]string{
		"operations",
		"ops",
	}
	for _, str := range keywords {
		if str == scanner.Text() {
			fmt.Println("<You may use the following operations:")
			for _, op := range ops_list {
				// ops_list is a global variable which may be seen a the beginning of this file
				if player.HasAbility[op[0]] {
					fmt.Println(op[0] + "\t" + op[1])
				}
			}
			// fmt.Println("<For more details on a given operation, enter \"details [operation]\" or \"d [operation]\". For example, for details on modus ponens enter \"details modus ponens\" or \"d mp\".")
		}
	}
}

// Checks the buffer for keywords related to user commands requesting details on logical operations
func CheckDetails(scanner *bufio.Scanner) {
	primary_keywords := [...]string{
		"details",
		"d",
	}
	for _, p := range primary_keywords {
		after, found := strings.CutPrefix(scanner.Text(), p)
		if found {
			trimmed := strings.Trim(after, " ")
			switch trimmed {
			case "mp", "modus ponens":
				fmt.Println("Modus Ponens may be used by entering \"modus ponens\" or \"mp\" followed by a letter indicating a statement and a letter indicating an implication.\nFor example:")
				fmt.Println(modus_ponens_example)
				// more cases here
			}
		}
	}
}

func HandleModusPonens(puz *Puzzle, input_stat string, input_imp string) []string {
	// Check if user chosen letters match puzzle letters for statement and implication
	stat_a, ok_stat_a := puz.Stats[strings.ToUpper(input_stat)]
	imp, ok_imp := puz.Imps[strings.ToUpper(input_imp)]
	stat_b := imp.Con
	_, ok_stat_b := puz.Stats[stat_b.Letter]

	fmt.Println(ok_stat_a, ok_stat_b, ok_imp)
	if ok_stat_a && ok_imp && ok_stat_b {
		// Apply modus ponens rule
		ModusPonens(&stat_a, &stat_b, &imp)
		// Update puzzle values
		puz.Stats[strings.ToUpper(input_stat)] = stat_a
		puz.Stats[imp.Con.Letter] = stat_b
		puz.Imps[strings.ToUpper(input_imp)] = imp
		return []string{
			"Modus Ponens applied!",
			stat_a.ToString(),
			imp.ToString(),
			" . ",
			". .",
			stat_b.ToString(),
		}
	}
	return []string{
		"Modus Ponens is not applicable to the selected symbols",
	}
}
