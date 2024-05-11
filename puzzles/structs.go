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
	IsHidden bool
}

func (s *Statement) Truth() *string {
	return &s.TruthVal
}

func (s *Statement) IsTrue() bool {
	return s.TruthVal == "true"
}

func (s *Statement) IsFalse() bool {
	return s.TruthVal == "false"
}

func (s *Statement) IsEmpty() bool {
	return s.TruthVal == "empty"
}

func (s *Statement) IsUnkown() bool {
	return s.TruthVal == "empty"
}

func (s *Statement) IsAnd() bool {
	if strings.Contains(s.Letter, "&") {
		str_array := strings.SplitAfter(s.Letter, "")
		if str_array[1] == "&" && str_array[0] != str_array[2] {
			return true
		}
	}
	if s.Letter[1] == '&' && len(s.Letter) == 3 {
		if s.Letter[0] != s.Letter[2] {
			return true
		}
	}
	return false
}

func (stat_a *Statement) TruthAndStat(stat_b *Statement) string {
	new_truthval := "This is an error"
	if stat_a.IsTrue() && stat_b.IsTrue() {
		new_truthval = "true"
	} else if stat_a.IsEmpty() || stat_b.IsEmpty() {
		new_truthval = "unknown"
	} else if stat_a.IsUnkown() || stat_b.IsUnkown() {
		new_truthval = "empty"
	} else {
		new_truthval = "false"
	}
	return new_truthval
}

// func (s *Statement) SplitAnd() ([]byte, bool) {
// 	if s.IsAnd() {
// 		return []byte{s.Letter[0], s.Letter[2]}, true
// 	}
// 	return []byte{}, false
// }
//

func (s *Statement) SplitAnd() ([]string, bool) {
	if s.IsAnd() {
		str_array := strings.SplitAfter(s.Letter, "")
		return []string{str_array[0], str_array[2]}, true
	}
	return []string{}, false
}

func (i *Implication) ContainsAndStats(stat_a1 string, stat_a2 string) bool {
	if i.Ant.Letter[0] == stat_a1[0] && i.Ant.Letter[2] == stat_a2[0] {
		return true
	} else if i.Ant.Letter[2] == stat_a1[0] && i.Ant.Letter[0] == stat_a2[0] {
		return true
	}
	return false
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

func (p *Puzzle) InsertAnd(stat_a *Statement, stat_b *Statement) {
	stat_letter_and := stat_a.Letter + "&" + stat_b.Letter
	stat_object_and := stat_a.Letter + " and " + stat_b.Letter

	new_truthval := stat_a.TruthAndStat(stat_b)
	new_and_stat := Statement{
		Letter:   stat_letter_and,
		IsNeg:    false,
		Subject:  "",
		Relation: "",
		Object:   stat_object_and,
		TruthVal: new_truthval,
	}
	p.Stats[stat_letter_and] = new_and_stat
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
			return true
		}
	}
	//  else {
	// 	fmt.Println("Modus Ponens does not apply in this case.")
	//    return false
	// }
	return false
}

// func ModusPonensAnd(stat_and *Statement, stat_b *Statement, imp *Implication) bool {
//   stat_split, ok_split := stat_and.SplitAnd()
//   if ok_split {
//     stat_a1 := stat_split[0]
//     stat_a2 := stat_split[1]
//     imp_stat_split, ok_split := imp.Ant.Letter.SplitAnd()
//     if ok_split {
//       imp_stat_a1
//     }
//   }
// 	if !stat_a1.IsEmpty() && !stat_a2.IsEmpty() {
// 			if imp_a_split == stat_a_split {
// 			}
// 			stat_b.TruthVal = stat_a.TruthVal
//
// 			return true
// 		}
// 	return false
// }

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

func ValidateAnd(stat_a *Statement, stat_b *Statement) bool { return true }
