package main

import (
	"github.com/rs/zerolog"
	"github.com/soulgarden/game/conf"
	"github.com/soulgarden/game/db"
	"github.com/soulgarden/game/entities"
	"github.com/soulgarden/game/services"
)

func main() {
	cfg := conf.New()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if cfg.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	pg := db.Connect(cfg)

	defer pg.Close()

	ws := &services.WorldService{}

	world := entities.NewWorld()
	ws.GenerateAreas(world)

	quit := make(chan int, 1)
	ws.Live(world, quit, cfg.TickInterval)
}
