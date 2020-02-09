package services

import (
	"math/rand"
	"time"

	"github.com/soulgarden/game/dictionaries"

	"github.com/rs/zerolog/log"
	"github.com/soulgarden/game/entities"
)

//WorldService handle all activities in world
type WorldService struct {
	playerService *PlayerService
}

//GenerateAreas generate nwe areas in new empty world
func (s *WorldService) GenerateAreas(w *entities.World) {
	i := 0

	for y := 0; y <= 100; y++ {
		for x := 0; x <= 100; x++ {
			w.AddArea(&entities.Area{
				X: int32(x),
				Y: int32(y),
				Z: 0,
			})
			i++
		}
	}

	log.Info().Int("areas total", i).Msg("areas were generated")
}

func (s *WorldService) addPlayer(p *entities.Player, w *entities.World) {
	w.AddPlayer(p)
}

func (s *WorldService) tick(w *entities.World) {
	//if s.isEventFired(dictionaries.PlayerBornChance) {
		s.addPlayer(entities.NewPlayer(dictionaries.RaceHuman), w)
	//}

	for _, player := range w.GetPlayers() {
		s.playerService.tick(player)
	}

	stat := w.GetStat()
	log.Info().Int("areas number", stat.AreasNumber).Int("players number", stat.PlayerNumber).Msg("world stat")
}

func (s *WorldService) isEventFired(chance uint8) bool {
	return rand.Intn(dictionaries.TotalPercent) < int(chance)
}

//Live start life in world
func (s *WorldService) Live(w *entities.World, q chan int, tickInterval uint8) {
	ticker := time.NewTicker(time.Second * time.Duration(tickInterval))
	i := 0

	for {
		select {
		case <-ticker.C:
			s.tick(w)
			i++
			log.Info().Int("tick number", i).Msg("tick")
		case <-q:
			log.Info().Msg("quit")
			return
		}
	}
}

//NewWorldService new world server instance
func NewWorldService(ps *PlayerService) *WorldService {
	return &WorldService{ps}
}
