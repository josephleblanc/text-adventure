package main

import "fmt"

// "golang.org/x/term"

func main() {
	all_inputs := initUserInput()
	input_map := mapUserInput(all_inputs)
	printTitle()

	helpInfo()

	promptEnter()

	printSlow("\tIt is late. Your eyes are growing bleary after staring at the screen for hours. Your posture, at first upright, is now hunched with exhaustion. You have been working on this text adventure for your CSCI-20 class for so long. But soon it will be complete.")
	centerText("\n\n<Type \"continue\" or \"c\" to continue, or type \"quit\" to exit the program>")
	promptContinue()
	printSlow("\tAs you hunt down the final bugs in your code, your eyes droop, your agile fingers begin to go limp, and you drowsily slump over as you fall into slumber, your cheek resting on your keyboard...\n")
	fmt.Printf("\n")
	elipsis()
	fmt.Printf("\n")
	// Introduce Aristotle
	printSlow("\tYou open your eyes to a stark white room with bare walls and no furniture. The only other occupant is an old man with a long white beard holding a golden sphere in his left hand. His eyes light up when he sees you and he introduces himself.\n")
	promptContinue()
	printSlow("\tHello sleeper, welcome to the dreamscape. Sometimes visitors arrive here after drifting off to sleep while thinking of logic, computing, or math. What is your name, sleeper?\n")
	player_data := promptName(input_map["name"])
	printSlow("Your name is " + player_data.name)
	printSlow("\tYou have arrived at a fortunate moment, as the dreamscape is currently dealing with an influx of irrationality thanks to the popularity of that damn Lewis Carrol.\n")
}
