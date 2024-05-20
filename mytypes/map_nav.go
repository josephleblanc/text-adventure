package mytypes

import (
	"fmt"
	"text-adventure/myprint"
)

func (m *Map) Travel(p *Player, d string) {
	room, exists := m.rooms[p.InRoom].GoesTo[d]
	if exists && room.is_locked {
		if len(p.Inventory) < 2 {
			myprint.PrintSlow(room.locked_description)
		} else if room.Name == "Final" && len(p.Inventory) == 2 {
			myprint.PrintSlow("With both keys clutched tightly in your hands, you approach the door to the north, its surface etched with cryptic symbols that seem to shift and writhe before your eyes. The locks themselves appear as twisted knots of metal, defying comprehension with their convoluted mechanisms. ")
			myprint.PrintSlow("As you fumble with the keys, the very air seems to thicken with uncertainty, as if reality itself is playing tricks on your senses. Each turn of the keys sends a jolt of disorientation coursing through your mind, as if the room is conspiring to confound your every move.")
			myprint.PrintSlow("With a final twist, the locks finally yield, but the door remains closed, its surface rippling like a mirage. For a moment, you question whether you've truly unlocked it or merely fallen deeper into the labyrinth of confusion. With a sense of trepidation, you push against the door, unsure of what lies on the other side, but determined to press forward regardless.")
			// m.UnlockRoom("Final")
			p.InRoom = room.Name
			move_descr, descr_exists := m.rooms[p.InRoom].MoveDescription[d]
			if descr_exists {
				myprint.PrintSlow(move_descr)
			}
			room.is_explored = true
			myprint.PrintSlow(room.Description)
		}
	} else {
		if exists {
			move_descr, descr_exists := m.rooms[p.InRoom].MoveDescription[d]
			p.InRoom = room.Name
			fmt.Println("In room", p.InRoom)
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
}

func (m *Map) Look(p *Player, direction string) {
	view, exists := m.rooms[p.InRoom].View[direction]
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

// func (m *Map) UnlockRoom(room_name string) {
// 	old_room, exists := m.rooms[room_name]
// 	if exists {
// 		new_room := old_room
// 		new_room.is_locked = false
// 		// m.rooms[room_name] = new_room
//
// 		// fmt.Println("New Room:")
// 		// fmt.Println(m.rooms[room_name])
//
// 		fmt.Println("Boss Room:")
// 		boss_room := Room{
// 			Name:               "Final",
// 			Description:        "You stand in the final room, its walls adorned with bewildering symbols and puzzles that defy comprehension. The air crackles with an unsettling energy, challenging you to unravel its mysteries and emerge victorious. This is the ultimate test of wit and intellect, the culmination of your journey through the realm of dreams.",
// 			GoesTo:             make(map[string]*Room),
// 			is_explored:        false,
// 			is_locked:          false,
// 			locked_description: "The door is heavy with two large locks, barring your way forward. Return with two keys if you wish to enter.",
// 		}
// 		fmt.Println(boss_room)
//
// 		m.rooms[room_name] = boss_room
//
// 	}
// }

func (m *Map) PlayerRoom(p *Player) (bool, string) {
	room, exists := m.rooms[p.InRoom]
	if exists {
		return true, room.Name
	} else {
		return false, ""
	}
}
