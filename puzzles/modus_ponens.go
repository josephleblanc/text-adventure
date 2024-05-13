package puzzles

import (
	"fmt"
	// "text-adventure/myprint"
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

	fmt.Println(stat_a.ToString() + "\n" + stat_b.ToString() + "\n" + imp_c.ToString() + "\n")
	myprint.PrintSlow("\tAristotle: We have here two statements (A) and (B), as well as an implication (C). We begin with the knowledge that (A) is true, and purpose to prove that (B) is true.")
	myprint.PrintSlow("\tAristotle: Why don't you try to use modus ponens to show that (B) is true? Try thinking \"modus ponens A C\", or if you prefer brevity, \"mp A C\"")
	fmt.Println()
	player.HasAbility["mp"] = true
	for puz.Stats["B"].TruthVal != "true" {
		// ^^ for loop contains win condition for puzzle
		PromptTool(&puz, &backup, player)
	}
}
