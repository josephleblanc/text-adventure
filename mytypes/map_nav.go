package mytypes

import (
	"text-adventure/myprint"
)

func (m *Map) Travel(p *Player, d string) {
	room, exists := m.rooms[p.InRoom].GoesTo[d]
	if exists && room.is_locked {
		myprint.PrintSlow(room.locked_description)
		_, room_name := m.PlayerRoom(p)
		if room_name == "Crossroads" && len(p.Inventory) == 2 {
			myprint.PrintSlow("With both keys clutched tightly in your hands, you approach the door to the north, its surface etched with cryptic symbols that seem to shift and writhe before your eyes. The locks themselves appear as twisted knots of metal, defying comprehension with their convoluted mechanisms. ")
			myprint.PrintSlow("As you fumble with the keys, the very air seems to thicken with uncertainty, as if reality itself is playing tricks on your senses. Each turn of the keys sends a jolt of disorientation coursing through your mind, as if the room is conspiring to confound your every move.")
			myprint.PrintSlow("With a final twist, the locks finally yield, but the door remains closed, its surface rippling like a mirage. For a moment, you question whether you've truly unlocked it or merely fallen deeper into the labyrinth of confusion. With a sense of trepidation, you push against the door, unsure of what lies on the other side, but determined to press forward regardless.")
			m.UnlockRoom("Final")
		}
		return
	}
	if exists {
		move_descr, descr_exists := m.rooms[p.InRoom].MoveDescription[d]
		p.InRoom = room.Name
		if !room.is_explored {
			if descr_exists {
				myprint.PrintSlow(move_descr)
			}
			room.is_explored = true
			myprint.PrintSlow(room.Description)
		} else {
			myprint.PrintSlow("You travel " + d + " to the " + room.Name + " room.")
		}
	} else {
		myprint.PrintSlow("There is no way to travel in that direction.")
	}
}

func (m *Map) Look(p *Player, d string) {
	view, exists := m.rooms[p.InRoom].View[d]
	if exists {
		myprint.PrintSlow(view)
	} else {
		myprint.PrintSlow("There is nothing of note in that direction.")
	}
}

func (m *Map) CurrentRoom(p *Player) {
	name := m.rooms[p.InRoom].Name
	if name != "" {
		to_print := "You are currently in the " + name + " room."
		myprint.PrintSlow(to_print)
	}
}

func (m *Map) IsPuzzleComplete(p *Player) bool {
	room, exists := m.rooms[p.InRoom]
	if exists {
		return room.PuzzleComplete
	} else {
		return false
	}
}

func (m *Map) CompletePuzzle(p *Player) {
	old_room, exists := m.rooms[p.InRoom]
	if exists {
		new_room := old_room
		new_room.PuzzleComplete = true
		m.rooms[p.InRoom] = new_room
	}
}

func (m *Map) UnlockRoom(room_name string) {
	old_room, exists := m.rooms[room_name]
	if exists {
		new_room := old_room
		new_room.is_locked = false
		m.rooms[room_name] = new_room
	}
}

func (m *Map) PlayerRoom(p *Player) (bool, string) {
	room, exists := m.rooms[p.InRoom]
	if exists {
		return true, room.Name
	} else {
		return false, ""
	}
}
