package puzzles

import (
	"fmt"
	"text-adventure/myprint"
	"text-adventure/mytypes"
)

func ModusPonensTutorial(player *mytypes.Player) {
	stat_a := Statement{
		Letter:   "A",
		IsNeg:    false,
		Subject:  "I",
		Relation: "am",
		Object:   "a man/woman",
		TruthVal: "true",
	}
	stat_b := Statement{
		Letter:   "B",
		IsNeg:    false,
		Subject:  "I",
		Relation: "am",
		Object:   "mortal",
		TruthVal: "unknown",
	}
	imp_c := ImpFrom("C", "true", &stat_a, &stat_b)

	stats := make(map[string]Statement)
	imps := make(map[string]Implication)

	stats["A"] = stat_a
	stats["B"] = stat_b
	imps["C"] = imp_c

	puz := Puzzle{
		Stats: stats,
		Imps:  imps,
	}
	backup := puz

	fmt.Println(stat_a.ToString() + "\n" + stat_b.ToString() + "\n" + imp_c.ToString())
	fmt.Println("\tAristotle: We have here two statements (A) and (B), as well as an implication (C). We begin with the knowledge that (A) is true, and purpose to prove that (B) is true.")
	fmt.Println("\tWhy don't you try to use modus ponens to show that (B) is true? Try thinking \"modus ponens A C\", or if you prefer brevity, \"mp A C\" ")
	player.HasAbility["mp"] = true
	for puz.Stats["B"].TruthVal != "true" {
		// ^^ for loop contains win condition for puzzle
		PromptTool(&puz, &backup, player)
	}
	myprint.PrintSlow("\nAristotle: Well done! I apologize if this is the first time you were made aware of this, but as a man/woman you are indeed mortal. This demonstrates the power and peril of symbolic logic - it allows us to arrive at new and sometimes uncomfortable truths.")
	myprint.PrintSlow("\nAristotle: Now that you have your first tool of logic, I can go back to thinking and leave all the work to you, my dear student. You may \"go\" wherever you wish in the four cardianl directions, just be sure to \"look\" and see if there is a passage there. Now, go clear the conundrums, and once you prove <win conition here>, you will be returned to your waking world.")

	// PrintSlow("A: " + first.words)
}
