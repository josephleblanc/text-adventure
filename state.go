package main

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

type playerData struct {
	name      string
	inventory []string
}

func newPlayerData(name string) playerData {
	p := playerData{name: name}
	p.inventory = []string{}
	return p
}

func (p *playerData) invent_add(item string) {
	p.inventory = append(p.inventory, item)
}
