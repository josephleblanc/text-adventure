package mytypes

type (
	Map struct {
		rooms map[string]Room
	}
	Room struct {
		Name        string
		Description string
		is_explored bool
		items       []Item
		npcs        []Npc
		objects     []Object
		// keys are directions, e.g. 'east', 'north'
		GoesTo             map[string]*Room
		MoveDescription    map[string]string
		View               map[string]string
		doors              []Door
		is_locked          bool
		locked_description string
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

// Interactable objects, e.g. light switch
type Object struct {
	name     string
	desc     string
	interact bool
}

type Npc struct {
	has_met     bool
	has_puzzle  bool
	has_truth   bool
	has_ability bool
	intro       string
	phrases     map[string]string
}

func (npc Npc) Dialogue(user_says string) {
	// todo
}

type Player struct {
	name   string
	truths []Statement
	// index of has_ability is Ability.name
	HasAbility     map[string]bool
	tutorial_done  bool
	aristotle_done bool
	in_challenge   bool
	can_move       bool
	// index of has_def is Definition.name
	has_defn map[string]bool
	// inventory []Item
	InRoom string
}

// Items the user can pick up and keep
type Item struct {
	name string
	desc string
	uses []string
}

// Abilities the user gains like 'contrapositive',
// or 'sum_convergent'
type Ability struct {
	name   string
	desc   string
	action func(stat Statement)
}

// Logical statements used in puzzles
type Statement struct {
	is_neg bool
	claim  string
}

// Useful definitions for the user to know when solving puzzles,
// e.g. 'logical proposition', 'convergent series'
type Definition struct {
	name string
	desc string
}
