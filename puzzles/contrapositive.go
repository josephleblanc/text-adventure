package puzzles

import (
	// "text-adventure/myprint"
	"fmt"
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
	for puz.Stats["A"].TruthVal != "true" {
		// ^^ for loop contains win condition for puzzle
		PromptTool(&puz, &backup, player)
	}
}
