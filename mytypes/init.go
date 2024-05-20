package mytypes

// Text Descriptions below written by ChatGPT3.
func InitMap() Map {
	start_room := Room{
		Name:        "Start",
		Description: "",
		View: map[string]string{
			"north": "To your north is a simple white door, made of an known substance. It lacks detail just as with the rest of this room, and you've no idea what may lie beyond.",
			"west":  "There is only a blank, white wall in this direction.",
			"east":  "There is only a blank, white wall in this direction.",
			"south": "There is only a blank, white wall in this direction.",
		},
		GoesTo:      make(map[string]*Room),
		is_explored: true,
	}
	crossroad := Room{
		// contains puzzle #2
		// Unlocking description:
		// With both keys clutched tightly in your hands, you approach the door to the north, its surface etched with cryptic symbols that seem to shift and writhe before your eyes. The locks themselves appear as twisted knots of metal, defying comprehension with their convoluted mechanisms.
		// As you fumble with the keys, the very air seems to thicken with uncertainty, as if reality itself is playing tricks on your senses. Each turn of the keys sends a jolt of disorientation coursing through your mind, as if the room is conspiring to confound your every move.
		// With a final twist, the locks finally yield, but the door remains closed, its surface rippling like a mirage. For a moment, you question whether you've truly unlocked it or merely fallen deeper into the labyrinth of confusion. With a sense of trepidation, you push against the door, unsure of what lies on the other side, but determined to press forward regardless.
		//
		// Go north after opening door:
		//You bravely step into the abyss beyond the open door, but instead of darkness, you find yourself in a disorienting whirlwind of shifting perspectives and surreal landscapes. The corridor twists and turns in impossible angles, defying all logic and reason.
		// The walls pulse with kaleidoscopic patterns, each glance revealing a different reality altogether. The air crackles with energy, sending shivers down your spine as you struggle to make sense of your surroundings.
		// Despite the chaos that surrounds you, a sense of purpose drives you forward, urging you to navigate this maze of confusion and unearth the truths hidden within its depths. With each bewildering step, you inch closer to the heart of the mystery that lies ahead.
		Name:        "Crossroads",
		Description: "You find yourself in a modest room adorned with shelves holding various kitchen utensils and cookware. A table sits in the center, upon which rests a perplexing puzzle laid out on a piece of parchment. The air is filled with the faint scent of spices and herbs, mingling with the faint crackle of a fire burning in the hearth. Amidst this cozy atmosphere, the puzzle's logical statements hang in the air, challenging your intellect and reasoning skills.",
		GoesTo:      make(map[string]*Room),
		is_explored: false,
		View: map[string]string{
			"north": "As you peer towards the northern expanse of the room, your attention is drawn to a sturdy door adorned with two imposing locks, standing as formidable barriers to further progress. Each lock glimmers with an aura of mystery, their mechanisms shrouded in enigma. The door itself seems to emanate a sense of anticipation, as if beckoning you to uncover the secrets that lie beyond its guarded threshold. With determination flickering in your heart, you know that unraveling the mysteries of nearby rooms will be the key to unlocking the path forward in the Land of Rationality.",
			"west":  "To the west, a door fashioned from polished mahogany catches your eye. Its surface is smooth and unadorned, save for a small brass handle that gleams softly in the ambient light.",
			"east":  "Your gaze turns to the east, where a door of wrought iron stands sentinel. Intricate scrollwork adorns its surface, intertwining in mesmerizing patterns that seem to shift and dance with every flicker of light.",
			"south": "Facing southward, you see a door crafted from ancient stone blocks, its surface weathered by the passage of time. A faint aura of solemnity surrounds it, hinting at the secrets it guards within.",
		},
		MoveDescription: map[string]string{
			"east": "You stride confidently towards the east, leaving behind the cozy ambiance of the puzzle-filled room. As you step through the doorway, you find yourself in a corridor bathed in soft, ambient light. The walls are adorned with intricate scrollwork, casting delicate shadows that dance across the polished marble floor. A sense of anticipation fills the air, urging you to press onward into the unknown depths of the Land of Rationality. With each step, you can't help but wonder what challenges and revelations await you around the next corner.",
			"west": "You proceed westward through the corridor, anticipation building with each step. As you reach the sturdy wooden door, you push it open, eager to uncover what lies beyond.",
		},
	}
	// puzzle # 15
	// win description:
	// As the final logical deduction falls into place, a faint shimmer dances across the parchment, drawing your attention. With a soft rustle, the parchment begins to unfurl, revealing a glint of metal nestled within its folds. As you reach out to investigate, your fingers brush against a small, intricately crafted key, its surface adorned with ornate patterns that seem to dance in the light.
	// The key feels weighty in your palm, its significance tangible as you realize its potential to unlock new pathways and hidden secrets within the realm of the Land of Rationality. With a sense of triumph and anticipation, you pocket the key, knowing that it will serve as a tangible symbol of your triumph over the puzzles that once confounded you.
	ducks := Room{
		Name:        "Duck",
		Description: "You enter a snug chamber, its walls adorned with rustic paintings depicting serene landscapes. A sturdy wooden table sits at the heart of the room, upon which rests a parchment bearing a perplexing puzzle. The faint aroma of ink mingles with the scent of freshly baked bread, creating a comforting atmosphere. As you contemplate the enigmatic statements before you, the room seems to resonate with the quacking of ducks and the dignified presence of officers, urging you to unlock the puzzle's secrets amidst this quaint setting.",
		GoesTo:      make(map[string]*Room),
		is_explored: false,
		View: map[string]string{
			"north": "Your gaze drifts towards the northern expanse of the room, revealing a charming window framed by billowing curtains. Beyond the glass, a tranquil scene unfolds, with lush greenery swaying gently in the breeze and sunlight dappling the landscape.",
			"west":  "Turning to the west, your eyes fall upon the sturdy wooden door through which you entered. Its weathered surface bears the marks of time, hinting at the countless journeys it has facilitated. A sense of familiarity washes over you as you contemplate the threshold leading back to the previous room.",
			"east":  "Casting your glance towards the eastern corner, you notice a cozy reading nook nestled amidst a collection of bookshelves. Soft cushions invite you to sink into their embrace, while the flickering light of a nearby lantern illuminates the titles of classic literature lining the shelves.",
			"south": "As you peer towards the southern boundary of the room, your attention is drawn to a charming fireplace adorned with a mantlepiece. The crackling flames cast a warm glow, casting dancing shadows across the room and filling the air with a comforting warmth.",
		},
	}
	// puzzle #19
	// win description:
	// As the puzzle is solved, a key materializes amidst the echoes in the room. It glows softly, its intricate design shimmering with significance. With a sense of triumph, you reach out and grasp it, feeling empowered to unlock new paths in the Land of Rationality.
	tapestry := Room{
		Name:        "Tapestry",
		Description: "You enter a room adorned with intricate tapestries, their patterns seeming to dance in anticipation of the puzzle's challenge. A table stands in the center, hosting a collection of logical statements that echo through the room. The air carries a subtle tension, mingling with the scent of mystery and the allure of unraveling the enigma. As you contemplate the statements, they speak of names and their suitability for heroes of romance, forming a labyrinth of logic that beckons you to explore its depths.",
		GoesTo:      make(map[string]*Room),
		is_explored: false,
		View: map[string]string{
			"north": "Your gaze turns northward, revealing a smooth expanse of wall adorned with intricate tapestries depicting scenes of logic and reason. The patterns seem to dance before your eyes, hinting at the depth of knowledge that lies within the confines of this room.",
			"west":  "Peering towards the west, you find a solid wall adorned with elegant motifs, its surface reflecting the ambient light in subtle hues. There are no apparent features to capture your attention in this direction, leaving your curiosity to linger elsewhere.",
			"east":  "As you glance towards the east, your eyes fall upon the sturdy wooden door through which you entered. Its weathered surface bears the marks of time and countless journeys, serving as a silent sentinel between the realms of the Land of Rationality.",
			"south": "Casting your gaze towards the southern boundary of the room, you find another expanse of wall adorned with intricate tapestries. The patterns seem to weave a tale of intellect and insight, beckoning you to delve deeper into the mysteries that lie ahead.",
		},
	}
	// puzzle #46
	// boss room
	// win description:
	// As you solve the final puzzle, a sense of clarity washes over you, dissolving the confusion and uncertainty that plagued your journey. With a deep breath, you emerge from the depths of the enigmatic room and find yourself back at your desk, the familiar surroundings grounding you in reality.
	// As you sit there, reflecting on the strange dream you just experienced, a subtle message catches your eye. A piece of parchment lies on the desk, bearing the unmistakable script of Aristotle himself. The message is cryptic, hinting at deeper meanings and hidden truths, leaving you to ponder whether the dream was merely a product of your imagination or a glimpse into a realm beyond comprehension.
	// With a sense of wonder and curiosity, you realize that the journey may have ended, but the mysteries of the Land of Rationality continue to linger in your mind. As you return to your daily life, Aristotle's final message serves as a reminder that the boundaries between reality and illusion are often more blurred than they seem.
	boss_room := Room{
		Name:               "Final",
		Description:        "You stand in the final room, its walls adorned with bewildering symbols and puzzles that defy comprehension. The air crackles with an unsettling energy, challenging you to unravel its mysteries and emerge victorious. This is the ultimate test of wit and intellect, the culmination of your journey through the realm of dreams.",
		GoesTo:             make(map[string]*Room),
		is_explored:        false,
		is_locked:          true,
		locked_description: "The door is heavy with two large locks, barring your way forward. Return with two keys if you wish to enter.",
	}

	crossroad.GoesTo["south"] = &start_room
	crossroad.GoesTo["east"] = &ducks
	crossroad.GoesTo["north"] = &boss_room
	crossroad.GoesTo["west"] = &tapestry

	tapestry.GoesTo["east"] = &crossroad

	ducks.GoesTo["west"] = &crossroad

	start_room.GoesTo["north"] = &crossroad

	boss_room.GoesTo["south"] = &crossroad

	new_map := map[string]Room{
		"Start":      start_room,
		"Crossroads": crossroad,
		"Duck":       ducks,
		"Tapestry":   tapestry,
		"Final":      boss_room,
	}
	return Map{rooms: new_map}
}
