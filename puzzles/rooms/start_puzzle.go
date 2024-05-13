package rooms

import (
	"text-adventure/mytypes"
)

func StartPuzzle(player *mytypes.Player) {
	switch player.InRoom {
	case "Crossroads":
		CrossroadsPuzzle(player)
	case "Duck":
		DuckPuzzle(player)
	case "Tapestry":
		TapestryPuzzle(player)
	case "Final":
		FinalPuzzle(player)
	}
}
