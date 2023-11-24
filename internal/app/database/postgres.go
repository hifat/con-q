package database

import (
	"fmt"

	"github.com/hifat/con-q-api/internal/app/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresConnection(cfg config.AppConfig) (*gorm.DB, func()) {
	dbCfg := cfg.DB
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbCfg.Host,
		dbCfg.Username,
		dbCfg.Password,
		dbCfg.Name,
		dbCfg.Port,
	)
	conn, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		panic("can't connect to database: " + err.Error())
	}

	cleanup := func() {
		sqlDB, err := conn.DB()
		if err == nil {
			sqlDB.Close()
		}
	}

	return conn, cleanup
}
