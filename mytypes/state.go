package mytypes

//// Sources
//
// usage of structs
// https://gobyexample.com/structs
// I've used their example code and made significant alterations to suite the
// program's purpse
//
// usage of empty slice
// https://gosamples.dev/empty-vs-nil-slice/
// I used the empty slice definition to help instantiate the struct
//
// usage of methods
// https://gobyexample.com/methods
// I've used their example code and made significant alterations to suite the
// program's purpse

type PlayerData struct {
	Name      string
	Inventory []string
}

func NewPlayerData(name string) PlayerData {
	p := PlayerData{Name: name}
	p.Inventory = []string{}
	return p
}

func (p *PlayerData) invent_add(item string) {
	p.Inventory = append(p.Inventory, item)
}
