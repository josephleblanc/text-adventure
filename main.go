package main

import (
	"fmt"
	"text-adventure/myprint"
	"text-adventure/mytypes"
	"text-adventure/puzzles"
	"text-adventure/puzzles/rooms"
	"text-adventure/utils"
)

func main() {
	// Input maps for handling user input such as name
	all_inputs := utils.InitUserInput()
	// Load text for dialogue and descriptions from file
	input_map := utils.MapUserInput(all_inputs)
	// Dialogue text stored in local file
	text_map := utils.CsvToTextList("text/intro.csv")
	// Player data that will change as the player gains access to the tools of logic
	player := mytypes.Player{
		HasAbility: make(map[string]bool),
		InRoom:     "Start",
		Inventory:  make(map[string]bool),
	}
	world_map := mytypes.InitMap()

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

	// Tutorial: modus ponens
	puzzles.ModusPonensTutorial(&player)
	myprint.PrintSlow("\tAristotle: Well done! I apologize if this is the first time you were made aware of this, but as a man/woman you are indeed mortal. This demonstrates the power and peril of symbolic logic - it allows us to arrive at new and sometimes uncomfortable truths.")
	utils.CenterText("<Type \"help\" or \"h\" at any time to review details on logic tools.>")
	fmt.Println()

	// Tutorial: Negation and Contrapositive
	myprint.PrintSlow("\tAristotle: You only need two more tools to solve all the puzzles ahead of you, Negation and Contrapositive. Negation allows you to invert a statement along with its truth value, while Contrapositive may be applied to an implication, such that the true statement \"If A then B\" becomes the true statement \"If not B then not A\"...")
	myprint.PrintSlow("\tAristotle: That may sound complicated, but I'm confident you'll pick up the trick soon enough, given the right motivation. Here, I'll manifest a some rain to encourage you. If you don't want to be rained on, just prove that it could not be raining - there can't be rain without clouds after all.")
	myprint.PrintSlow("\tJust as Aristotle finishes speaking, you begin to feel drops of water falling on your skin. There is indeed rain falling on you! This is rather bewildering, as the rain is not falling from clouds but through the ceiling above.")
	myprint.PrintSlow("\tAristotle: Now, to use contrapositive, just think \"contrapositive c\" or \"cp c\", and to use negation, just think \"negation b\" or \"neg b\" ")
	puzzles.ContraPositiveTutorial(&player)

	myprint.PrintSlow("\tAristotle: Well done!")
	fmt.Println()
	// Conclude Aristotle scene
	myprint.PrintSlow("\tAristotle: Now that you have your logic tools, I can go back to thinking and leave all the work to you, my dear student. You may \"go\" wherever you wish in the four cardinal directions (north, west, south, east), just be sure to \"look\" and see if there is a passage there. Now, go clear the conundrums, and once you complete the puzzle beyond the locked door in the next room, you will be returned to your waking world.")

	// Main game loop
	final_complete := false
	for !final_complete {
		utils.PromptNav(&player, &world_map)
		if !world_map.IsPuzzleComplete(&player) {
			switch _, room_name := world_map.PlayerRoom(&player); room_name {
			case "Crossroads":
				rooms.CrossroadsPuzzle(&player)
			case "Duck":
				rooms.DuckPuzzle(&player)
			case "Tapestry":
				rooms.TapestryPuzzle(&player)
			case "Final":
				rooms.FinalPuzzle(&player)
				final_complete = true
			}
			world_map.CompletePuzzle(&player)
		}
	}

	// Wrap up game and run credits
	myprint.PrintSlow("\tAs you solve the final puzzle, a sense of clarity washes over you, dissolving the confusion and uncertainty that plagued your journey. With a deep breath, you emerge from the depths of the enigmatic room and find yourself back at your desk, the familiar surroundings grounding you in reality.")
	myprint.PrintSlow("\tAs you sit there, reflecting on the strange dream you just experienced, a subtle message catches your eye. A piece of parchment lies on the desk, bearing the unmistakable script of Aristotle himself. The message is cryptic, hinting at deeper meanings and hidden truths, leaving you to ponder whether the dream was merely a product of your imagination or a glimpse into a realm beyond comprehension. ")
	myprint.PrintSlow("\tWith a sense of wonder and curiosity, you realize that the journey may have ended, but the mysteries of the Land of Rationality continue to linger in your mind. As you return to your daily life, Aristotle's final message serves as a reminder that the boundaries between reality and illusion are often more blurred than they seem.")
	utils.Ellipsis()
	myprint.PrintSlow("\tThank you for playing this logic game!")
	utils.Ellipsis()
	myprint.PrintSlow("\tMade by Brasides")
}
