package puzzles

import (
	"fmt"
	"strings"
)

type Statement struct {
	Letter   string
	IsNeg    bool
	Subject  string
	Relation string
	Object   string
	TruthVal string
}

type Puzzle struct {
	Stats map[string]Statement
	Imps  map[string]Implication
}

func (p *Puzzle) Status() {
	for _, stat := range p.Stats {
		fmt.Println(stat.ToString())
	}
	for _, imp := range p.Imps {
		fmt.Println(imp.ToString())
	}
}

func (s *Statement) NegString() string {
	neg_string := ""
	if s.IsNeg {
		neg_string = " not"
	}
	return s.Relation + neg_string
}

func (s *Statement) ToString() string {
	truth_string := "?:    "
	switch s.TruthVal {
	case "true":
		truth_string = "True: "
	case "false":
		truth_string = "False:"
	}
	all_strings := []string{
		"(" + s.Letter + ")\t\t",
		truth_string,
		s.Subject,
		s.NegString(),
		s.Object,
	}
	return strings.Join(all_strings, " ")
}

type Implication struct {
	// Symbolic letter
	Letter string
	// Truth Value
	TruthVal string
	// Antecedent
	Ant Statement
	// Consequent
	Con Statement
}

func ImpFrom(letter string, truth_val string, ant *Statement, con *Statement) Implication {
	ant_copy := *ant
	ant_copy.TruthVal = "empty"
	con_copy := *con
	con_copy.TruthVal = "empty"
	return Implication{
		Letter:   letter,
		TruthVal: truth_val,
		Ant:      ant_copy,
		Con:      con_copy,
	}
}

func (i *Implication) ToString() string {
	symbols := "(" + i.Letter + ": " + i.Ant.Letter + "->" + i.Con.Letter + ")\t"
	truth_string := "?:    "
	switch i.TruthVal {
	case "true":
		truth_string = "True: "
	case "false":
		truth_string = "False:"
	}
	all_strings := []string{
		symbols,
		truth_string,
		"If",
		i.Ant.Subject,
		i.Ant.NegString(),
		i.Ant.Object + ",",
		"then",
		i.Con.Subject,
		i.Con.NegString(),
		i.Con.Object + ".",
	}
	return strings.Join(all_strings, " ")
}

func ModusPonens(stat_a *Statement, stat_b *Statement, imp *Implication) {
	if stat_a.TruthVal == "true" || stat_a.TruthVal == "false" {
		if *imp == ImpFrom(imp.Letter, imp.TruthVal, stat_a, stat_b) {
			stat_b.TruthVal = stat_a.TruthVal
			fmt.Println(stat_a.ToString())
			fmt.Println(imp.ToString())
			fmt.Println(stat_b)
		}
	} else {
		fmt.Println("Modus Ponens does not apply in this case.")
	}
}
