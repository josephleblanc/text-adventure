package main

type (
	Map struct {
		room []Room
	}
	Room struct {
		is_explored bool
		items       []Item
		npcs        []Npc
		objects     []Object
		// keys are directions, e.g. 'east', 'north'
		goes_to []*Room
		doors   []Door
	}
	Door struct {
		is_open bool
		goes_to *Room
		key     string
	}
)

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
	has_ability    map[string]bool
	tutorial_done  bool
	aristotle_done bool
	in_challenge   bool
	can_move       bool
	// index of has_def is Definition.name
	has_defn  map[string]bool
	inventory []Item
	in_room   Room
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
