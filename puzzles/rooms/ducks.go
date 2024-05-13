package rooms

import (
	// "text-adventure/myprint"

	"text-adventure/myprint"
	"text-adventure/mytypes"
	"text-adventure/puzzles"
)

// 15.
//
//	(1) All ducks in this village, that are branded ‘B,’ belong to Mrs. Bond;
//	(2) Ducks in this village never wear lace collars, unless they are branded ‘B’;
//	(3) Mrs. Bond has no gray ducks in this village.
//
// Univ. “ducks in this village”; a = belonging to Mrs. Bond; b = branded ‘B’; c = gray; d = wearing lace-collars.
//
// Solution:
//
//	No gray ducks in this village wear lace collars.
//
// TODO: Add a puzzle problem statement to each puzzle
func DuckPuzzle(player *mytypes.Player) {
	stat_a := puzzles.Statement{
		Letter:   "A",
		IsNeg:    false,
		Subject:  "All ducks in this village, that",
		Relation: "are",
		Object:   "branded 'B', belong to Mrs. Bond",
		TruthVal: "true",
	}

	stat_b := puzzles.Statement{
		Letter:   "B",
		IsNeg:    false,
		Subject:  "Ducks in this village never",
		Relation: "wear",
		Object:   "lace collars, unless they are branded 'B'",
		TruthVal: "true",
	}

	stat_c := puzzles.Statement{
		Letter:  "C",
		IsNeg:   false,
		Subject: "Mrs. Bond",
		// TODO: Add a way to change the negation word when the relation is "has"
		Relation: "has",
		Object:   "no gray ducks in this village",
		TruthVal: "true",
	}

	stat_d := puzzles.Statement{
		Letter:   "D",
		IsNeg:    false,
		Subject:  "No ducks in this village",
		Relation: "are",
		Object:   "branded 'B'",
		TruthVal: "unknown",
	}

	stat_e := puzzles.Statement{
		Letter:   "E",
		IsNeg:    false,
		Subject:  "No gray ducks in this village",
		Relation: "wear",
		Object:   "lace collars",
		TruthVal: "unknown",
	}

	stat_a_c := puzzles.Statement{
		Letter:   "A&C",
		IsNeg:    false,
		Subject:  "",
		Relation: "",
		Object:   "A and C (test)",
		TruthVal: "unknown",
		// IsHidden: true,
	}
	stat_b_d := puzzles.Statement{
		Letter:   "B&D",
		IsNeg:    false,
		Subject:  "",
		Relation: "",
		Object:   "B and D (test)",
		TruthVal: "unknown",
		// IsHidden: true,
	}

	stat_d_copy := stat_d
	// stat_d_copy.Negate()
	imp_f := puzzles.ImpFrom("F", "true", &stat_a_c, &stat_d_copy)
	imp_g := puzzles.ImpFrom("G", "true", &stat_b_d, &stat_e)

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

	for puz.Stats["E"].TruthVal != "true" {
		// ^^ for loop contains win condition for puzzle
		puzzles.PromptTool(&puz, &backup, player)
	}
	// TODO: Add flavor text here
	myprint.PrintSlow("As the final logical deduction falls into place, a faint shimmer dances across the parchment, drawing your attention. With a soft rustle, the parchment begins to unfurl, revealing a glint of metal nestled within its folds. As you reach out to investigate, your fingers brush against a small, intricately crafted key, its surface adorned with ornate patterns that seem to dance in the light. The key feels weighty in your palm, its significance tangible as you realize its potential to unlock new pathways and hidden secrets within the realm of the Land of Rationality. With a sense of triumph and anticipation, you pocket the key, knowing that it will serve as a tangible symbol of your triumph over the puzzles that once confounded you.")
	myprint.PrintSlow("You now have the key from this room, and are one step closer to unlocking the final door.")
	myprint.PrintSlow("\tHaving solved the puzzle, you are free to \"look\" around the room or \"go\" in a direction.")
}
