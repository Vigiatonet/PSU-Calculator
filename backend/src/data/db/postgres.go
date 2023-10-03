package db

import (
	"fmt"
	"time"

	"github.com/Vigiatonet/PSU-Calculator/config"
	"github.com/Vigiatonet/PSU-Calculator/pkg/logging"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var psqlDB *gorm.DB
var logger = logging.NewLogger(config.GetConfig())

func InitDB(cfg *config.Config) (err error) {
	cnn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.DbName,
	)
	psqlDB, err = gorm.Open(postgres.Open(cnn), &gorm.Config{})
	if err != nil {
		return err
	}
	db, err := psqlDB.DB()
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	db.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
	db.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	db.SetConnMaxLifetime(cfg.Postgres.ConnMaxLifetime * time.Minute)
	logger.Info(logging.Postgres, logging.Startup, "postgres started", nil)
	return nil
}

func GetDB() *gorm.DB {
	return psqlDB
}

func CloseDB() {
	db, err := psqlDB.DB()
	if err != nil {
		logger.Fatal(err, logging.Postgres, logging.SubCategory(logging.Postgres), "cant Get postgresDB", nil)
	}
	err = db.Close()
	if err != nil {
		logger.Fatal(err, logging.Postgres, logging.SubCategory(logging.Postgres), "cant close postgresDB", nil)
	}
}
