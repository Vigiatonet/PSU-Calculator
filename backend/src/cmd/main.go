package main

import (
	"github.com/Vigiatonet/PSU-Calculator/api"
	"github.com/Vigiatonet/PSU-Calculator/config"
	"github.com/Vigiatonet/PSU-Calculator/data/cache"
	"github.com/Vigiatonet/PSU-Calculator/data/db"
	"github.com/Vigiatonet/PSU-Calculator/data/db/migrations"
	"github.com/Vigiatonet/PSU-Calculator/pkg/logging"
)

func main() {
	cfg := config.GetConfig()
	log := logging.NewLogger(cfg)

	err := db.InitDB(cfg)
	if err != nil {
		log.Fatal(err, logging.Internal, logging.Startup, "cant init db", nil)
	}
	err = cache.LoadRedis(cfg)
	if err != nil {
		log.Fatal(err, logging.Redis, logging.Startup, "cant init redisDB", nil)
	}
	migrations.Up_01()
	api.InitServer(cfg)

}
