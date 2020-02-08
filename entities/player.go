package entities

import "github.com/soulgarden/game/dictionaries"

//Player entity
type Player struct {
	Age           int32
	health        int32
	physicalPower int32
	magicPower    int32
	coins         int32
	isAlive       bool
	luckiness     int32
	hunger        int32
	Location      struct {
		X int32
		Y int32
		Z int32
	}
	equipment struct {
		weapon int32
		head   int32
		armor  int32
		shoes  int32
	}
	Race int32
}

//NewPlayer nwe player instance
func NewPlayer(race int32) *Player {
	return &Player{
		Age:           dictionaries.StartAge,
		health:        0,
		physicalPower: 0,
		magicPower:    0,
		coins:         0,
		isAlive:       false,
		luckiness:     0,
		hunger:        0,
		Location: struct {
			X int32
			Y int32
			Z int32
		}{
			X: 0,
			Y: 0,
			Z: 0,
		},
		equipment: struct {
			weapon int32
			head   int32
			armor  int32
			shoes  int32
		}{
			weapon: 0,
			head:   0,
			armor:  0,
			shoes:  0,
		},
		Race: race,
	}
}
