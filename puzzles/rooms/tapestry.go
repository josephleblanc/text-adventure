package rooms

import (
	// "text-adventure/myprint"

	"text-adventure/myprint"
	"text-adventure/mytypes"
	"text-adventure/puzzles"
)

// 19.
//
//	(1) No name in this list is unsuitable for the hero of a romance;
//	(2) Names beginning with a vowel are always melodious;
//	(3) No name is suitable for the hero of a romance, if it begins with a consonant.
//
// Univ. “names”; a = beginning with a vowel; b = in this list; c = melodious; d = suitable for the hero of a romance.
//
// Solution:
//
//	No name in this list is unmelodious.
//
// TODO: Add a puzzle problem statement to each puzzle
func TapestryPuzzle(player *mytypes.Player) {
	stat_a := puzzles.Statement{
		Letter:   "A",
		IsNeg:    false,
		Subject:  "No name in this list",
		Relation: "is",
		Object:   "unsuitable for the hero of a romance",
		TruthVal: "true",
	}

	stat_b := puzzles.Statement{
		Letter:   "B",
		IsNeg:    false,
		Subject:  "Names beginning with a vowel",
		Relation: "are",
		Object:   "always melodious",
		TruthVal: "true",
	}

	stat_c := puzzles.Statement{
		Letter:  "C",
		IsNeg:   false,
		Subject: "No name",
		// TODO: Add a way to change the negation word when the relation is "has"
		Relation: "is",
		Object:   "suitable for the hero of a romance, if it begins with a consonant",
		TruthVal: "true",
	}

	stat_d := puzzles.Statement{
		Letter:   "D",
		IsNeg:    false,
		Subject:  "No name in this list",
		Relation: "is",
		Object:   "begins with a vowel",
		TruthVal: "unknown",
	}

	stat_e := puzzles.Statement{
		Letter:   "E",
		IsNeg:    false,
		Subject:  "No name in this list",
		Relation: "is",
		Object:   "unmelodious",
		TruthVal: "unknown",
	}

	stat_a_c := puzzles.Statement{
		Letter:   "A&C",
		IsNeg:    false,
		Subject:  "",
		Relation: "",
		Object:   "A and C",
		TruthVal: "unknown",
		IsHidden: true,
	}
	stat_b_d := puzzles.Statement{
		Letter:   "B&D",
		IsNeg:    false,
		Subject:  "",
		Relation: "",
		Object:   "B and D",
		TruthVal: "unknown",
		IsHidden: true,
	}

	stat_d_copy := stat_d
	stat_d_copy.Negate()
	imp_f := puzzles.ImpFrom("F", "true", &stat_a_c, &stat_d_copy)
	imp_f.Con.TruthVal = "true"
	imp_f.Con.IsNeg = false
	imp_g := puzzles.ImpFrom("G", "true", &stat_b_d, &stat_e)
	imp_g.Con.TruthVal = "true"
	imp_g.Con.IsNeg = false

	stats := make(map[string]puzzles.Statement)
	imps := make(map[string]puzzles.Implication)

	stats["A"] = stat_a
	stats["B"] = stat_b
	stats["C"] = stat_c
	stats["D"] = stat_d
	stats["E"] = stat_e
	//
	// stats["A&C"] = stat_d
	// stats["D&E"] = stat_e

	imps["F"] = imp_f
	imps["G"] = imp_g

	puz := puzzles.Puzzle{
		Stats: stats,
		Imps:  imps,
	}
	backup := puz

	puz.Status()

	// TODO: put flavor text here
	myprint.PrintSlow("Find the truth of all statements to solve the puzzle!")
	for puz.Stats["E"].TruthVal != "true" {
		// ^^ for loop contains win condition for puzzle
		puzzles.PromptTool(&puz, &backup, player)
	}
	myprint.PrintSlow("\tAs the puzzle is solved, a key materializes amidst the echoes in the room. It glows softly, its intricate design shimmering with significance. With a sense of triumph, you reach out and grasp it, feeling empowered to unlock new paths in the Land of Rationality.")
	myprint.PrintSlow("\tYou have taken one more step on the path to unlocking the final door... who knows what lays beyond?")
	player.Inventory["tapestry_key"] = true
}
