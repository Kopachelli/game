package entities

import (
	"sync"
)

//World world
type World struct {
	mx sync.RWMutex

	areas   []*Area
	players []*Player
}

//Area in world
type Area struct {
	X    int32
	Y    int32
	Z    int32
	Type int32
}

//WorldStat represent world statistic
type WorldStat struct {
	AreasNumber  int
	PlayerNumber int
}

//AddArea add new area to world
func (w *World) AddArea(a *Area) {
	w.mx.Lock()
	w.areas = append(w.areas, a)
	w.mx.Unlock()
}

//AddPlayer add new player to world
func (w *World) AddPlayer(p *Player) {
	w.mx.Lock()
	w.players = append(w.players, p)
	w.mx.Unlock()
}

//GetPlayers get players in world
func (w *World) GetPlayers() []*Player {
	return w.players
}

//GetStat get world statistic
func (w *World) GetStat() *WorldStat {
	return &WorldStat{
		AreasNumber:  len(w.areas),
		PlayerNumber: len(w.players),
	}
}

//NewWorld new world instance
func NewWorld() *World {
	return &World{}
}
