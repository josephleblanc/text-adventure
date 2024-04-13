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

func (s *Statement) Truth() *string {
	return &s.TruthVal
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

func (s *Statement) Negate() {
	s.IsNeg = !s.IsNeg
	if s.TruthVal == "true" {
		s.TruthVal = "false"
	}
	if s.TruthVal == "false" {
		s.TruthVal = "true"
	}
}

func (s *Statement) ToString() string {
	truth_string := "?:    "
	switch s.TruthVal {
	case "true":
		truth_string = "True: "
	case "false":
		truth_string = "False:"
	}
	neg_symb := ""
	if s.IsNeg {
		neg_symb = "!"
	}
	all_strings := []string{
		"(" + neg_symb + s.Letter + ")\t\t",
		truth_string,
		s.Subject,
		s.NegString(),
		s.Object,
	}
	return strings.Join(all_strings, " ")
}

// type HasTruth interface {
// 	Truth() *string
// 	IsNegInterface() *bool
// }

// func (s *Statement) IsNegInterface() *bool {
// 	return &s.IsNeg
// }
// func (s *Implication) IsNegInterface() *bool {
// 	return &s.IsNeg
// }

// func NegSelf[T HasTruth](t T) bool {
// 	*t.IsNegInterface() = !*t.IsNegInterface()
// 	if *t.Truth() == "true" {
// 		*t.Truth() = "false"
// 		return true
// 	} else if *t.Truth() == "false" {
// 		*t.Truth() = ""
// 		return true
// 	}
// 	return false
// }

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

func (i *Implication) Truth() *string {
	return &i.TruthVal
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
	ant_neg_symb := ""
	con_neg_symb := ""
	if i.Ant.IsNeg {
		ant_neg_symb = "!"
	}
	if i.Con.IsNeg {
		con_neg_symb = "!"
	}
	symbols := "(" + i.Letter + ": " + ant_neg_symb + i.Ant.Letter + "->" + con_neg_symb + i.Con.Letter + ")\t"
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

// Applies the modus ponens transformation to an implication, if the antecedent
// has a known truth value.
func ModusPonens(stat_a *Statement, stat_b *Statement, imp *Implication) bool {
	if stat_a.TruthVal == "true" || stat_a.TruthVal == "false" {
		if *imp == ImpFrom(imp.Letter, imp.TruthVal, stat_a, stat_b) {
			stat_b.TruthVal = stat_a.TruthVal
			fmt.Println(stat_a.ToString())
			fmt.Println(imp.ToString())
			fmt.Println(stat_b)
			return true
		}
	}
	//  else {
	// 	fmt.Println("Modus Ponens does not apply in this case.")
	//    return false
	// }
	return false
}

func ContraPositive(imp *Implication) bool {
	hold_ant := imp.Ant
	imp.Ant = imp.Con
	imp.Con = hold_ant
	imp.Ant.Negate()
	imp.Con.Negate()
	fmt.Println(imp.ToString())
	return true
}

func Negate(stat *Statement) bool {
	if stat.TruthVal == "true" {
		stat.TruthVal = "false"
	} else if stat.TruthVal == "false" {
		stat.TruthVal = "true"
	} else {
		return false
	}
	stat.IsNeg = !stat.IsNeg
	return true
}
