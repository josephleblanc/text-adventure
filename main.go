package main

import (
	"fmt"
	"text-adventure/myprint"
	"text-adventure/mytypes"
	"text-adventure/puzzles"
	"text-adventure/puzzles/rooms"
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
		InRoom:     "Start",
	}
	world_map := mytypes.InitMap()

	rooms.FinalPuzzle(&player)

	// puzzles.ModusPonensTutorial(&player)

	// player.HasAbility["mp"] = true
	// player.HasAbility["cp"] = true
	// puzzles.ContraPositiveTutorial(&player)

	printTitle()

	utils.HelpInfo()

	utils.PromptEnter()

	myprint.PrintSlow("\t" + text_map["intro1"])
	// utils.CenterText("\n\n<Type \"continue\" or \"c\" to continue, or type \"quit\" to exit the program>")
	// utils.PromptContinue()
	myprint.PrintSlow("\t" + text_map["intro2"])
	utils.Ellipsis()
	fmt.Println()
	fmt.Println()
	// Introduce Aristotle
	myprint.PrintSlow("\t" + text_map["aristotle1"])
	myprint.PrintSlow("\t" + text_map["aristotle2"])
	player_data := utils.PromptName(input_map["name"])
	myprint.PrintSlow("\t" + text_map["aristotle3.1"] + player_data.Name + text_map["aristotle3.2"])
	myprint.PrintSlow("\t" + text_map["aristotle4"])
	myprint.PrintSlow("\t" + text_map["aristotle5"])
	myprint.PrintSlow("\t" + text_map["aristotle6"])
	fmt.Println()

	// Tutorial: modus ponens
	puzzles.ModusPonensTutorial(&player)
	myprint.PrintSlow("\tAristotle: Well done! I apologize if this is the first time you were made aware of this, but as a man/woman you are indeed mortal. This demonstrates the power and peril of symbolic logic - it allows us to arrive at new and sometimes uncomfortable truths.")
	// TODO: Add a way for the user to check which logic tools they have access to, along with their descriptions.
	utils.CenterText("<Type \"help\" or \"h\" at any time to review details on logic tools.>")
	fmt.Println()

	// Tutorial: Negation and Contrapositive
	myprint.PrintSlow("\tAristotle: You only need two more tools to solve all the puzzles ahead of you, Negation and Contrapositive. Negation allows you to invert a statement along with its truth value, while Contrapositive may be applied to an implication, such that the true statement \"If A then B\" becomes the true statement \"If not B then not A\"...")
	myprint.PrintSlow("\tAristotle: That may sound complicated, but I'm confident you'll pick up the trick soon enough, given the right motivation. Here, I'll manifest a some rain to encourage you. If you don't want to be rained on, just prove that it could not be raining - there can't be rain without clouds after all.")
	myprint.PrintSlow("\tJust as Aristotle finishes speaking, you begin to feel drops of water falling on your skin. There is indeed rain falling on you! This is rather bewildering, as the rain is not falling from clouds but through the ceiling above.")
	puzzles.ContraPositiveTutorial(&player)

	myprint.PrintSlow("\tAristotle: Well done!")
	fmt.Println()
	// Conclude Aristotle scene
	// TODO: Add win condition in the dialogue below:
	myprint.PrintSlow("\tAristotle: Now that you have your logic tools, I can go back to thinking and leave all the work to you, my dear student. You may \"go\" wherever you wish in the four cardinal directions (north, west, south, east), just be sure to \"look\" and see if there is a passage there. Now, go clear the conundrums, and once you prove <win conition here>, you will be returned to your waking world.")
	for {
		utils.PromptNav(&player, &world_map)
	}
}
