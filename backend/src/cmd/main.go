package main

import (
	"github.com/Vigiatonet/PSU-Calculator/src/api"
	"github.com/Vigiatonet/PSU-Calculator/src/config"
	"github.com/Vigiatonet/PSU-Calculator/src/data/cache"
	"github.com/Vigiatonet/PSU-Calculator/src/data/db"
	"github.com/Vigiatonet/PSU-Calculator/src/pkg/logging"
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
	api.InitServer(cfg)
}
