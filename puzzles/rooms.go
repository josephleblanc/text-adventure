package puzzles

import (
	// "text-adventure/myprint"
	"fmt"
	"text-adventure/myprint"
	"text-adventure/mytypes"
)

// Puzzle for the crossroads room
// This is a Lewis Carroll logic puzzle, taken from "Symbolic Logic", by Lewis Carrol
// Puzzle #2, found in Section 9:
// https://www.gutenberg.org/cache/epub/28696/pg28696-images.html#p070dagger
//
//	(1) My saucepans are the only things I have that are made of tin;
//	(2) I find all your presents very useful;
//	(3) None of my saucepans are of the slightest use.
//
// Univ. “things of mine”; a = made of tin; b = my saucepans; c = useful; d = your presents.
//
// Solution:
//
//	Your presents to me are not made of tin.
//
// Worked Solution:
//
//	Made of tin (!(saucepan complement))
//	AND Your presents (!(useful complement))
//	AND my saucepands (!useful);
//
//	complement(saucepans + made of tin)
//	AND saucepans useful
//	AND complement( your presents useful )
//
//	Therefore:
//	!(made of tin + your presents)
//	AND your presents
//
//	Therefore
//	Your presents + !made of tin
func CrossroadsPuzzle(player *mytypes.Player) {
	stat_a := Statement{
		Letter:   "A",
		IsNeg:    false,
		Subject:  "My saucepans",
		Relation: "are",
		Object:   "the only things I have that are made of tin",
		TruthVal: "true",
	}

	stat_b := Statement{
		Letter:   "B",
		IsNeg:    false,
		Subject:  "I",
		Relation: "find",
		Object:   "all your presents very useful",
		TruthVal: "true",
	}

	stat_c := Statement{
		Letter:   "C",
		IsNeg:    false,
		Subject:  "None of my saucepans",
		Relation: "are",
		Object:   "of the smallest use",
		TruthVal: "true",
	}

	stat_d := Statement{
		Letter:   "D",
		IsNeg:    false,
		Subject:  "Things made of tin",
		Relation: "are",
		Object:   "useful",
		TruthVal: "unknown",
	}

	stat_e := Statement{
		Letter:   "E",
		IsNeg:    true,
		Subject:  "Your presents",
		Relation: "are",
		Object:   "made of tin",
		TruthVal: "unknown",
	}

	stat_a_c := Statement{
		Letter:   "A&C",
		IsNeg:    false,
		Subject:  "",
		Relation: "",
		Object:   "",
		TruthVal: "unknown",
	}
	stat_b_d := Statement{
		Letter:   "B&D",
		IsNeg:    false,
		Subject:  "",
		Relation: "",
		Object:   "",
		TruthVal: "unknown",
	}

	stat_d_copy := stat_d
	stat_d_copy.Negate()
	imp_f := ImpFrom("F", "true", &stat_a_c, &stat_d_copy)
	imp_g := ImpFrom("G", "true", &stat_b_d, &stat_e)

	stats := make(map[string]Statement)
	imps := make(map[string]Implication)

	stats["A"] = stat_a
	stats["B"] = stat_b
	stats["C"] = stat_c
	stats["D"] = stat_d
	stats["E"] = stat_e

	stats["A&C"] = stat_d
	stats["D&E"] = stat_e

	imps["F"] = imp_f
	imps["G"] = imp_g

	puz := Puzzle{
		Stats: stats,
		Imps:  imps,
	}
	backup := puz

	fmt.Println(stat_a.ToString())
	fmt.Println(stat_b.ToString())
	fmt.Println(stat_c.ToString())
	fmt.Println(stat_d.ToString())
	fmt.Println(stat_e.ToString())

	// TODO: Add text introducing the "and" ability
	player.HasAbility["and"] = true
	for puz.Stats["E"].TruthVal != "true" {
		// ^^ for loop contains win condition for puzzle
		PromptTool(&puz, &backup, player)
	}
	// TODO: Add flavor text here
	myprint.PrintSlow("\t<Say something at the end of the puzzle>")
}

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
func duck_puzzle(player *mytypes.Player) {
	stat_a := Statement{
		Letter:   "A",
		IsNeg:    false,
		Subject:  "My saucepans",
		Relation: "are",
		Object:   "the only things I have that are made of tin",
		TruthVal: "true",
	}

	stat_b := Statement{
		Letter:   "B",
		IsNeg:    false,
		Subject:  "I",
		Relation: "find",
		Object:   "all your presents very useful",
		TruthVal: "true",
	}

	stat_c := Statement{
		Letter:   "C",
		IsNeg:    false,
		Subject:  "None of my saucepans",
		Relation: "are",
		Object:   "of the smallest use",
		TruthVal: "true",
	}

	stat_d := Statement{
		Letter:   "D",
		IsNeg:    false,
		Subject:  "Things made of tin",
		Relation: "are",
		Object:   "useful",
		TruthVal: "unknown",
	}

	stat_e := Statement{
		Letter:   "E",
		IsNeg:    true,
		Subject:  "Your presents",
		Relation: "are",
		Object:   "made of tin",
		TruthVal: "unknown",
	}

	stat_a_c := Statement{
		Letter:   "A&C",
		IsNeg:    false,
		Subject:  "",
		Relation: "",
		Object:   "",
		TruthVal: "unknown",
	}
	stat_b_d := Statement{
		Letter:   "B&D",
		IsNeg:    false,
		Subject:  "",
		Relation: "",
		Object:   "",
		TruthVal: "unknown",
	}

	stat_d_copy := stat_d
	stat_d_copy.Negate()
	imp_f := ImpFrom("F", "true", &stat_a_c, &stat_d_copy)
	imp_g := ImpFrom("G", "true", &stat_b_d, &stat_e)

	stats := make(map[string]Statement)
	imps := make(map[string]Implication)

	stats["A"] = stat_a
	stats["B"] = stat_b
	stats["C"] = stat_c
	stats["D"] = stat_d
	stats["E"] = stat_e

	stats["A&C"] = stat_d
	stats["D&E"] = stat_e

	imps["F"] = imp_f
	imps["G"] = imp_g

	puz := Puzzle{
		Stats: stats,
		Imps:  imps,
	}
	backup := puz

	fmt.Println(stat_a.ToString())
	fmt.Println(stat_b.ToString())
	fmt.Println(stat_c.ToString())
	fmt.Println(stat_d.ToString())
	fmt.Println(stat_e.ToString())

	// TODO: Add text introducing the "and" ability
	player.HasAbility["and"] = true
	for puz.Stats["E"].TruthVal != "true" {
		// ^^ for loop contains win condition for puzzle
		PromptTool(&puz, &backup, player)
	}
	// TODO: Add flavor text here
	myprint.PrintSlow("\t<Say something at the end of the puzzle>")
}

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
func tapestry_puzzle(player *mytypes.Player) {
}

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
func final_puzzle(player *mytypes.Player) {
}
