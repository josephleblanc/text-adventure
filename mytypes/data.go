package mytypes

type (
	Map struct {
		rooms map[string]Room
	}
	Room struct {
		Name        string
		Description string
		is_explored bool
		// items       []Item
		// npcs        []Npc
		// objects     []Object
		// keys are directions, e.g. 'east', 'north'
		GoesTo             map[string]*Room
		MoveDescription    map[string]string
		View               map[string]string
		doors              []Door
		is_locked          bool
		locked_description string
		HasPuzzle          bool
		PuzzleComplete     bool
	}
	Door struct {
		is_open bool
		goes_to *Room
		key     string
	}
)

// A type of user input that
type UserInput struct {
	// The value of the user input
	Selection string
	// The first part of the sentence to verify input
	// e.g. "Your name is "
	PreVerify string
	// The second part of the sentence to verify input
	// e.g. ", is that right?"
	PostVerify string
	// The descriptive name of the value the user is asked to input.
	// e.g. "name"
	ToSelect string
	// The field title presented to the user.
	// e.g. "Name: "
	FieldName string
}

type Player struct {
	name string
	// index of has_ability is Ability.name
	HasAbility     map[string]bool
	tutorial_done  bool
	aristotle_done bool
	in_challenge   bool
	// can_move       bool
	// index of has_def is Definition.name
	has_defn  map[string]bool
	InRoom    string
	Inventory map[string]bool
}

// Abilities the user gains like 'contrapositive',
// or 'sum_convergent'
type Ability struct {
	name string
	desc string
}

// Useful definitions for the user to know when solving puzzles,
// e.g. 'logical proposition', 'convergent series'
type Definition struct {
	name string
	desc string
}
