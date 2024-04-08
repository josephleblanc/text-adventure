package mytypes

func InitMap() map[string]Room {
	start_room := Room{
		Name:        "start",
		description: "",
		is_explored: true,
	}
	negation_room := Room{
		Name:        "negation",
		description: "As you enter the large room you notice that it is mostly empty, save for a series of pillars spread unevenly through the room. Suddenly, one of the pillars disappears with a small \"pop\". A few moments later another appears just as suddenly.",
		is_explored: false,
	}
	start_goes_to := map[string]*Room{
		"north": &negation_room,
	}
	negation_goes_to := map[string]*Room{
		"south": &start_room,
	}

	start_room.GoesTo = start_goes_to
	negation_room.GoesTo = negation_goes_to

	new_map := map[string]Room{
		"start":    start_room,
		"negation": negation_room,
	}
	return new_map
}
