package services

import (
	"github.com/rs/zerolog/log"
	"github.com/soulgarden/game/dictionaries"
	"github.com/soulgarden/game/entities"
	"math/rand"
)

//PlayerService
type PlayerService struct{}

func (s *PlayerService) tick(p *entities.Player) {
	s.move(p)
}

func (s *PlayerService) move(p *entities.Player) {
	p.Location.X = rand.Intn(dictionaries.MaxX)
	p.Location.Y = rand.Intn(dictionaries.MaxY)

	log.Info().Str("ID", p.ID).Int("x", p.Location.X).Int("y", p.Location.Y)
}

//NewPlayerService new player service instance
func NewPlayerService() *PlayerService {
	return &PlayerService{}
}
