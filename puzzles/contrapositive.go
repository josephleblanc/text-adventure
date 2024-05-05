package puzzles

import (
	// "text-adventure/myprint"
	"fmt"
	"text-adventure/myprint"
	"text-adventure/mytypes"
)

func ContraPositiveTutorial(player *mytypes.Player) {
	stat_a := Statement{
		Letter:   "A",
		IsNeg:    false,
		Subject:  "It",
		Relation: "is",
		Object:   "raining",
		TruthVal: "unknown",
	}
	stat_b := Statement{
		Letter:   "B",
		IsNeg:    false,
		Subject:  "It",
		Relation: "is",
		Object:   "cloudy",
		TruthVal: "false",
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

	player.HasAbility["cp"] = true
	player.HasAbility["neg"] = true
	for puz.Stats["A"].TruthVal != "true" {
		// ^^ for loop contains win condition for puzzle
		// TODO: Add flavor text for the rain continuing to fall.
		// Possibly from Chatgpt3. If so, add an option to change text color for gpt text.
		PromptTool(&puz, &backup, player)
	}
	myprint.PrintSlow("\tThe moment you provide incontravertable proof that it must not be raining, the rain immediately stops. You are, however, still wet.")
}
