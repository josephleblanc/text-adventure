package main

import (
	"fmt"
	"text-adventure/myprint"
	"text-adventure/mytypes"
	"text-adventure/puzzles"
	"text-adventure/utils"
)

// "golang.org/x/term"

func main() {
	all_inputs := utils.InitUserInput()
	// Input maps for handling user input such as name
	input_map := utils.MapUserInput(all_inputs)
	// Load text for dialogue and descriptions from file
	text_map := utils.CsvToTextList("text/intro.csv")
	// Player data that will change as the player gains access to the tools of logic
	player := mytypes.Player{
		HasAbility: make(map[string]bool),
	}
	puzzles.ContraPositiveTutorial(&player)

	printTitle()

	utils.HelpInfo()

	utils.PromptEnter()

	myprint.PrintSlow("\t" + text_map["intro1"])
	utils.CenterText("\n\n<Type \"continue\" or \"c\" to continue, or type \"quit\" to exit the program>")
	utils.PromptContinue()
	myprint.PrintSlow("\t" + text_map["intro2"])
	fmt.Printf("\n")
	utils.Ellipsis()
	fmt.Printf("\n")
	// Introduce Aristotle
	myprint.PrintSlow("\t" + text_map["aristotle1"])
	utils.PromptContinue()
	myprint.PrintSlow("\t" + text_map["aristotle2"])
	player_data := utils.PromptName(input_map["name"])
	myprint.PrintSlow("\t" + text_map["aristotle3.1"] + player_data.Name + text_map["aristotle3.2"])
	utils.PromptContinue()
	myprint.PrintSlow("\t" + text_map["aristotle4"])
	utils.PromptContinue()
	myprint.PrintSlow("\t" + text_map["aristotle5"])
	utils.PromptContinue()
	myprint.PrintSlow("\t" + text_map["aristotle6"])

	puzzles.ModusPonensTutorial(&player)
	myprint.PrintSlow("\tAristotle: Well done! I apologize if this is the first time you were made aware of this, but as a man/woman you are indeed mortal. This demonstrates the power and peril of symbolic logic - it allows us to arrive at new and sometimes uncomfortable truths.")
	// TODO: Add a way for the user to check which logic tools they have access to, along with their descriptions.
	myprint.PrintSlow("\tAristotle: If you ever forget the logic tools, just think \"help\" or \"h\"")

	// Conclude Aristotle scene
	myprint.PrintSlow("\tAristotle: Now that you have your first tool of logic, I can go back to thinking and leave all the work to you, my dear student. You may \"go\" wherever you wish in the four cardinal directions (north, west, south, east), just be sure to \"look\" and see if there is a passage there. Now, go clear the conundrums, and once you prove <win conition here>, you will be returned to your waking world.")
}
