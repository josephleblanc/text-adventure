package main

import (
	"fmt"
	"text-adventure/utils"
)

// "golang.org/x/term"

func main() {
	all_inputs := utils.InitUserInput()
	input_map := utils.MapUserInput(all_inputs)
	printTitle()

	utils.HelpInfo()

	utils.PromptEnter()

	text_map := utils.CsvToTextList("text/intro.csv")

	utils.PrintSlow("\t" + text_map["intro1"])
	utils.CenterText("\n\n<Type \"continue\" or \"c\" to continue, or type \"quit\" to exit the program>")
	utils.PromptContinue()
	utils.PrintSlow("\t" + text_map["intro2"])
	fmt.Printf("\n")
	utils.Ellipsis()
	fmt.Printf("\n")
	// Introduce Aristotle
	utils.PrintSlow("\t" + text_map["aristotle1"])
	utils.PromptContinue()
	utils.PrintSlow("\t" + text_map["aristotle2"])
	player_data := utils.PromptName(input_map["name"])
	utils.PrintSlow("\t" + text_map["aristotle3.1"] + player_data.Name + text_map["aristotle3.2"])
}
