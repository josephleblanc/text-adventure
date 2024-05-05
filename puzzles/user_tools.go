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

var modus_ponens_example string = `
(A)		 True:  I am a man/woman
(B)		 ?:     I am mortal
(C: A->B)	 True:  If I am a man/woman, then I am mortal.

---------> modus ponens a c <---------
        --OR--
---------> mp a c <-------------------

(A)		 True:  I am a man/woman
(C: A->B)	 True:  If I am a man/woman, then I am mortal.
 .
. .
(B)		 True:  I am mortal`

var contrapositive_example string = `
(A)		 ?:     It is raining
(B)		 False: It is cloudy
(C: A->B)	 True:  If It is raining, then It is cloudy.

---------> contrapositive c <---------
        --OR--
---------> cp c <---------------------

(C: !B->!A)	 True:  If It is not cloudy, then It is not raining.
Contra-Positive Applied!!
(C: A->B)	 True:  If It is raining, then It is cloudy.
 .
. .
(C: !B->!A)	 True:  If It is not cloudy, then It is not raining.
`

var negation_example string = `
(A)		 ?:     It is raining
(B)		 False: It is cloudy
(C: A->B)	 True:  If It is raining, then It is cloudy.

---------> negation b <---------
        --OR--
---------> neg b <--------------

Negation Applied!
(B)		 False: It is cloudy
 .
. .
(!B)		 True:  It is not cloudy
`

// A complete list of operations with a short description of each
var ops_list [][]string = [][]string{
	{"mp", "modus ponens: Given statement A true and the implication A->B true, therefore B true"},
	{"cp", "contrapositive: Take the negation and reverse positions of statements in an implication. For example, the true implication, A->B, apply contrapositive to get the true implication !B->!A."},
	{"neg", "negation: If a statement is true, then the opposite of that statement is false. For example, if A is true, !A is false. Conversely, if !B is true, B is false."},
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

	// check for modus ponens command (long version)
	if len(split) == 4 && (split[0] == "modus" && split[1] == "ponens") {
		vals := HandleModusPonens(puz, split[2], split[3])
		for _, v := range vals {
			fmt.Println(v)
		}
		// check for modus ponens command (short version)
	} else if len(split) == 3 && split[0] == "mp" {
		vals := HandleModusPonens(puz, split[1], split[2])
		for _, v := range vals {
			fmt.Println(v)
		}
		// check for contrapositive command
	} else if len(split) == 2 && (split[0] == "cp" || split[1] == "contra-positive") {
		vals := HandleContraPositive(puz, split[1])
		for _, v := range vals {
			fmt.Println(v)
		}
	} else if len(split) == 2 && split[0] == "neg" {
		vals := HandleNegation(puz, split[1])
		for _, v := range vals {
			fmt.Println(v)
		}
	} else if len(split) == 1 && (split[0] == "status") {
		puz.Status()
	} else if len(split) == 1 && split[0] == "" {
		PuzzleHelp()
	}
}

func PuzzleHelp() {
	utils.CenterText("<You are currently in a puzzle.>")
	fmt.Println("Try applying the tools you have to solve the puzzle, or you may type \"status\" to see the current status of the puzzle and its statements, \"operations\" or \"ops\" to see available operations, or \"definitions\" or \"defs\" to see collected definitions.")
	fmt.Println("To reset the puzzle, type \"reset\"")
	fmt.Println()
}

func CheckReset(puzzle *Puzzle, backup *Puzzle, scanner *bufio.Scanner) {
	if scanner.Text() == "reset" {
		*puzzle = *backup
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
			fmt.Println("You may use the following operations:")
			for _, op := range ops_list {
				// ops_list is a global variable which may be seen a the beginning of this file
				if player.HasAbility[op[0]] {
					fmt.Println(op[0] + "\t" + op[1])
				}
			}
			fmt.Println("For more details on a given operation, enter \"details [operation]\" or \"d [operation]\". For example, for details on modus ponens enter \"details modus ponens\" or \"d mp\".")
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
				utils.CenterText("Modus Ponens")
				fmt.Println("Modus Ponens may be used by entering \"modus ponens\" or \"mp\" followed by a letter indicating a statement and a letter indicating an implication.\nFor example:")
				fmt.Println(modus_ponens_example + "\n")
			case "cp", "contrapositive":
				utils.CenterText("Contrapositive")
				fmt.Println("Contrapositive may be used by entering \"cp\" or \"contrapositive\" followed by a letter indicating an implication.\nFor example:")
				fmt.Println(contrapositive_example + "\n")
			case "neg", "negation":
				utils.CenterText("Negation")
				fmt.Println("Negation may be used by entering \"neg\" or \"negation\" followed by a letter indicating a statement. Then statement will be made false and the claim inverted.\nFor example:")
				fmt.Println(negation_example + "\n")
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

	if ok_stat_a && ok_imp && ok_stat_b {
		// Apply modus ponens rule
		is_applied := ModusPonens(&stat_a, &stat_b, &imp)
		if is_applied {
			// Update puzzle values
			puz.Stats[strings.ToUpper(input_stat)] = stat_a
			puz.Stats[imp.Con.Letter] = stat_b
			puz.Imps[strings.ToUpper(input_imp)] = imp
			// Return lines to print for user
			return []string{
				"Modus Ponens applied!",
				stat_a.ToString(),
				imp.ToString(),
				" . ",
				". .",
				stat_b.ToString(),
				"\n",
			}

		}
		return []string{
			"Modus Ponens is not applicable to the selected symbols. For Modus Ponens to be applicable, the truth value of the consequent must be known. Enter \"help mp\" for more details.",
		}
	}
	return []string{
		"Invalid symbols used. Please use the symbols from the puzzle, or enter \"h\" for help",
	}
}

func HandleContraPositive(puz *Puzzle, input_imp string) []string {
	// Verify user has input an implication
	imp, ok_imp := puz.Imps[strings.ToUpper(input_imp)]
	if ok_imp {
		old_imp_string := imp.ToString()
		is_applied := ContraPositive(&imp)
		if is_applied {
			// Update puzzle values
			puz.Imps[strings.ToUpper(input_imp)] = imp

			// Return lines to print for user
			return []string{
				"Contra-Positive Applied!!",
				old_imp_string,
				" . ",
				". .",
				imp.ToString(),
				"\n",
			}
		}

	}
	return []string{
		"Contra-Positive is not applicable to the selected symbol",
	}
}

func HandleNegation(puz *Puzzle, input_stat string) []string {
	stat, ok_stat := puz.Stats[strings.ToUpper(input_stat)]
	if ok_stat {
		old_in_string := stat.ToString()
		is_applied := Negate(&stat)
		if is_applied {

			// Update puzzle values
			puz.Stats[strings.ToUpper(input_stat)] = stat
			// NegSelf(&stat)
			new_string := stat.ToString()
			return []string{
				"Negation Applied!",
				old_in_string,
				" . ",
				". .",
				new_string,
				"\n",
			}
		}
	}
	return []string{
		"Negation cannot be applied to the symbol entered.",
	}
}
