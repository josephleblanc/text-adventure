package rooms

import (
	// "text-adventure/myprint"

	"text-adventure/myprint"
	"text-adventure/mytypes"
	"text-adventure/puzzles"
)

// 46.
//
//	(1) When I work a Logic-example without grumbling, you may be sure it is one that I can understand;
//	(2) These Soriteses are not arranged in regular order, like the examples I am used to;
//	(3) No easy example ever make my head ache;
//	(4) I ca’n’t understand examples that are not arranged in regular order, like those I am used to;
//	(5) I never grumble at an example, unless it gives me a headache.
//
// Univ. “Logic-examples worked by me”; a = arranged in regular order, like the examples I am used to; b = easy; c = grumbled at by me; d = making my head ache; e = these Soriteses; h = understood by me.
//
// Solution:
//
//	These Sorites-examples are difficult.

// TODO: Add a puzzle problem statement to each puzzle
func FinalPuzzle(player *mytypes.Player) {
	stat_a := puzzles.Statement{
		Letter:   "A",
		IsNeg:    false,
		Subject:  "When I work a Logic-example without grumbling, you may be sure it",
		Relation: "is",
		Object:   "one that I can understand",
		TruthVal: "true",
	}

	stat_b := puzzles.Statement{
		Letter:   "B",
		IsNeg:    false,
		Subject:  "These Soriteses",
		Relation: "are",
		Object:   "arranged in regular order, like the examples I am used to",
		TruthVal: "true",
	}

	stat_c := puzzles.Statement{
		Letter:  "C",
		IsNeg:   false,
		Subject: "No easy example ever",
		// TODO: Add a way to change the negation word when the relation is "made"
		Relation: "made",
		Object:   "my head ache",
		TruthVal: "true",
	}

	stat_d := puzzles.Statement{
		Letter:   "D",
		IsNeg:    true,
		Subject:  "I can't understand examples that",
		Relation: "are",
		Object:   "arranged in regular order, like those I am used to",
		TruthVal: "false",
	}

	stat_e := puzzles.Statement{
		Letter: "E",
		// TODO: Fix the negation for this one, as it comes out pretty awkwardly.
		IsNeg:    false,
		Subject:  "I never grumble at an example, unless it",
		Relation: "gives",
		Object:   "me a headache",
		TruthVal: "true",
	}

	stat_f := puzzles.Statement{
		Letter:   "F",
		IsNeg:    true,
		Subject:  "I",
		Relation: "am",
		Object:   "grumbling",
		TruthVal: "unknown",
		IsHidden: false,
	}

	stat_g := puzzles.Statement{
		Letter:   "G",
		IsNeg:    false,
		Subject:  "You may be sure it is a logic puzzle I",
		Relation: "do",
		Object:   "understand",
		TruthVal: "unknown",
		IsHidden: false,
	}

	// Implication !F -> G
	// If I am grumbling, then you may be sure it is a logic puzzle I do not understand
	stat_f_copy_fg := stat_f
	stat_f_copy_fg.TruthVal = "false"
	stat_g_copy_fg := stat_g
	stat_g_copy_fg.TruthVal = "true"
	// TODO: Test this
	imp_w := puzzles.ImpFrom("W", "true", &stat_f_copy_fg, &stat_g_copy_fg)
	// imp_fg.Con.TruthVal = "true"
	// imp_fg.Con.IsNeg = false

	// If B&D -> !G
	// If
	//	(B) These logic puzzles are not arranged in regular order, like the examples I am used to;
	//	AND
	//	(D) I can’t understand logic puzzles that are not arranged in regular order,
	//	    like those I am used to;
	//	--> !G: I can not understand this puzzle
	stat_bd := puzzles.Statement{
		Letter:   "B&D",
		IsNeg:    false,
		Subject:  "",
		Relation: "",
		Object:   "B and D",
		TruthVal: "unknown",
		IsHidden: true,
	}

	stat_bd_copy := stat_bd
	stat_bd_copy.TruthVal = "true"
	stat_bd_copy.IsHidden = false
	stat_g_copy := stat_g
	stat_g_copy.IsNeg = true
	stat_g_copy.TruthVal = "true"
	// TODO: Test this
	imp_x := puzzles.ImpFrom("X", "true", &stat_bd_copy, &stat_g_copy)

	stat_h := puzzles.Statement{
		Letter:   "H",
		IsNeg:    false,
		Subject:  "My head",
		Relation: "is",
		Object:   "aching",
		TruthVal: "unknown",
		IsHidden: false,
	}
	stat_ef := puzzles.Statement{
		Letter:   "E&F",
		IsNeg:    false,
		Subject:  "",
		Relation: "",
		Object:   "E and F",
		TruthVal: "unknown",
		IsHidden: true,
	}

	stat_ef_copy := stat_ef
	stat_ef_copy.TruthVal = "true"
	stat_ef_copy.IsHidden = false

	stat_h_copy := stat_h
	stat_h_copy.TruthVal = "true"
	// TODO: Test this
	imp_y := puzzles.ImpFrom("Y", "true", &stat_ef_copy, &stat_h_copy)

	stat_ch := puzzles.Statement{
		Letter:   "C&H",
		IsNeg:    false,
		Subject:  "",
		Relation: "",
		Object:   "C and H",
		TruthVal: "unknown",
		IsHidden: true,
	}

	stat_i := puzzles.Statement{
		Letter:   "I",
		IsNeg:    true,
		Subject:  "This puzzle",
		Relation: "is",
		Object:   "easy",
		TruthVal: "unknown",
		IsHidden: false,
	}

	stat_ch_copy := stat_ch
	stat_ch_copy.TruthVal = "true"
	stat_i_copy := stat_i
	stat_i_copy.TruthVal = "true"
	// TODO: Test this
	imp_z := puzzles.ImpFrom("Z", "true", &stat_ch_copy, &stat_i_copy)

	stats := make(map[string]puzzles.Statement)
	imps := make(map[string]puzzles.Implication)

	stats["A"] = stat_a
	stats["B"] = stat_b
	stats["C"] = stat_c
	stats["D"] = stat_d
	stats["E"] = stat_e
	stats["F"] = stat_f
	stats["G"] = stat_g
	stats["H"] = stat_h
	stats["I"] = stat_i
	//
	// stats["B&D"] = stat_d
	// stats["E&F"] = stat_e
	// stats["C&H"] = stat_e

	imps["W"] = imp_w
	imps["X"] = imp_x
	imps["Y"] = imp_y
	imps["Z"] = imp_z

	puz := puzzles.Puzzle{
		Stats: stats,
		Imps:  imps,
	}
	backup := puz

	puz.Status()

	myprint.PrintSlow("Prove that this is not an easy puzzle to complete your adventure in the dream!")
	for puz.Stats["I"].TruthVal != "true" {
		// ^^ for loop contains win condition for puzzle
		puzzles.PromptTool(&puz, &backup, player)
	}
	// TODO: Add flavor text here
	myprint.PrintSlow("\t<Say something at the end of the puzzle>")
}
