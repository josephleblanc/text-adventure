package mytypes

import "text-adventure/myprint"

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
			myprint.PrintSlow(room.description)
		} else {
			myprint.PrintSlow("You travel " + d + " to the " + room.Name + "room.")
		}
	} else {
		myprint.PrintSlow("There is no way to travel in that direction.")
	}
}
