package mytypes

import "text-adventure/myprint"

// func (p *Puzzle) Status() {
func (m *Map) Travel(p *Player, d string) {
	room, exists := m.rooms[p.InRoom].GoesTo[d]
	if exists {
		p.InRoom = room.Name
		if !room.is_explored {
			room.is_explored = true
			myprint.PrintSlow(room.description)
		} else {
			myprint.PrintSlow("You travel " + d + " to the " + room.Name + "room.")
		}
	} else {
		myprint.PrintSlow("There is no way to travel in that direction.")
	}
}
