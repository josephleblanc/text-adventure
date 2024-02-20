package main

import "fmt"

// "golang.org/x/term"

func main() {
	print_title()

	center_text("<Press \"Enter\" to continue, or type \"quit\" to exit the program>")
	promptEnter()

	printSlow("\tIt is late. Your eyes are growing bleary after staring at the screen for hours. Your posture, at first upright, is now hunched with exhaustion. You have been working on this text adventure for your CSCI-20 class for so long. But soon it will be complete.")
	center_text("\n\n<Press \"Enter\" to continue, or type \"quit\" to exit the program>")
	promptEnter()
	printSlow("\tAs you hunt down the final bugs in your code, your eyes droop, your agile fingers begin to go limp, and you drowsily slump over as you fall into slumber, your cheek resting on your keyboard...\n")
	fmt.Printf("\n")
	elipsis()
	fmt.Printf("\n")
}
