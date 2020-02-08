package db

import (
	"log"
	"strconv"

	"github.com/go-pg/pg/v9"
	"github.com/soulgarden/game/conf"
)

//Connect connect to DB
func Connect(cfg *conf.Config) *pg.DB {
	cp := pg.Connect(&pg.Options{
		User:     cfg.DB.User,
		Database: cfg.DB.Name,
		Addr:     cfg.DB.Host + ":" + strconv.Itoa(cfg.DB.Port),
		Password: cfg.DB.Password,
		PoolSize: cfg.DB.PoolSize,
	})

	ping(cp)

	return cp
}

func ping(cp *pg.DB) {
	if _, err := cp.Exec("SELECT 1"); err != nil {
		log.Fatal("PGSQL is down ", err)
	}
}
