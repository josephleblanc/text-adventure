package mytypes

import (
	"text-adventure/myprint"
)

func (m *Map) Travel(p *Player, d string) {
	room, exists := m.rooms[p.InRoom].GoesTo[d]
	if exists && room.is_locked {
		myprint.PrintSlow(room.locked_description)
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

func (m *Map) PlayerRoom(p *Player) (bool, string) {
	room, exists := m.rooms[p.InRoom]
	if exists {
		return true, room.Name
	} else {
		return false, ""
	}
}
