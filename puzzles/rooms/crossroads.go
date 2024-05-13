package rooms

import (
	// "text-adventure/myprint"

	"fmt"
	"text-adventure/myprint"
	"text-adventure/mytypes"
	"text-adventure/puzzles"
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
	stat_a := puzzles.Statement{
		Letter:   "A",
		IsNeg:    false,
		Subject:  "My saucepans",
		Relation: "are",
		Object:   "the only things I have that are made of tin",
		TruthVal: "true",
	}

	stat_b := puzzles.Statement{
		Letter:   "B",
		IsNeg:    false,
		Subject:  "I",
		Relation: "find",
		Object:   "all your presents very useful",
		TruthVal: "true",
	}

	stat_c := puzzles.Statement{
		Letter:   "C",
		IsNeg:    false,
		Subject:  "None of my saucepans",
		Relation: "are",
		Object:   "of the smallest use",
		TruthVal: "true",
	}

	stat_d := puzzles.Statement{
		Letter:   "D",
		IsNeg:    false,
		Subject:  "Things made of tin",
		Relation: "are",
		Object:   "useful",
		TruthVal: "unknown",
	}

	stat_e := puzzles.Statement{
		Letter:   "E",
		IsNeg:    true,
		Subject:  "Your presents",
		Relation: "are",
		Object:   "made of tin",
		TruthVal: "unknown",
	}

	stat_a_c := puzzles.Statement{
		Letter:   "A&C",
		IsNeg:    false,
		Subject:  "",
		Relation: "",
		Object:   "A and C",
		TruthVal: "empty",
		IsHidden: true,
	}
	stat_b_d := puzzles.Statement{
		Letter:   "B&D",
		IsNeg:    false,
		Subject:  "",
		Relation: "",
		Object:   "B and D",
		TruthVal: "empty",
		IsHidden: true,
	}

	stat_d_copy := stat_d
	stat_d_copy.Negate()
	imp_f := puzzles.ImpFrom("F", "true", &stat_a_c, &stat_d_copy)
	imp_g := puzzles.ImpFrom("G", "true", &stat_b_d, &stat_e)

	stats := make(map[string]puzzles.Statement)
	imps := make(map[string]puzzles.Implication)

	stats["A"] = stat_a
	stats["B"] = stat_b
	stats["C"] = stat_c
	stats["D"] = stat_d
	stats["E"] = stat_e

	stats["A&C"] = stat_a_c
	stats["D&E"] = stat_b_d

	imps["F"] = imp_f
	imps["G"] = imp_g

	puz := puzzles.Puzzle{
		Stats: stats,
		Imps:  imps,
	}
	backup := puz

	// fmt.Println(stat_a.ToString())
	// fmt.Println(stat_b.ToString())
	// fmt.Println(stat_c.ToString())
	// fmt.Println(stat_d.ToString())
	// fmt.Println(stat_e.ToString())

	puz.Status()
	fmt.Println()
	// TODO: Add help information for the "and" tool
	myprint.PrintSlow("Just as the puzzle comes into being, and you wonder how you can solve it with Aristotle's tools, you hear his disembodied voice sound through the room:")
	myprint.PrintSlow("\tAristotle: Oh! I almost forgot, there is one more tool you will need - the \"and\" tool. You may use it by saying \"and a c\". I would explain more, but I'm sure you'll figure it out!")
	myprint.PrintSlow("\tAristotle: The main thing to remember with \"and\" statements is that you can use them with modus ponens by thinking \"mp a&c f\", once you have constructed them.")
	player.HasAbility["and"] = true
	for puz.Stats["E"].TruthVal != "true" {
		// ^^ for loop contains win condition for puzzle
		puzzles.PromptTool(&puz, &backup, player)
	}
	myprint.PrintSlow("As the final logical deduction falls into place, a faint shimmer dances across the parchment, drawing your attention. With a soft rustle, the parchment begins to unfurl, revealing a glint of metal nestled within its folds. As you reach out to investigate, your fingers brush against a small, intricately crafted key, its surface adorned with ornate patterns that seem to dance in the light. The key feels weighty in your palm, its significance tangible as you realize its potential to unlock new pathways and hidden secrets within the realm of the Land of Rationality. With a sense of triumph and anticipation, you pocket the key, knowing that it will serve as a tangible symbol of your triumph over the puzzles that once confounded you.")
	myprint.PrintSlow("You now have the key from this room, and are one step closer to unlocking the final door.")
	myprint.PrintSlow("\tHaving solved the puzzle, you are free to \"look\" around the room or \"go\" in a direction.")
}
