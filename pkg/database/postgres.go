package database

import (
	"database/sql"
	"fmt"
	"github.com/ahmadrezamusthafa/ecommerce/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	Timezone = "Asia/Jakarta"
	SSL      = "disable"
)

func NewPostgresqlDatabase(cfg config.DatabaseConfiguration) (*gorm.DB, error) {
	dsn := getDSN(cfg)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.MaxOpen)
	db.SetMaxIdleConns(cfg.MaxIdle)
	db.SetConnMaxLifetime(cfg.MaxLifetime)

	orm, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	return orm, err
}

func getDSN(cfg config.DatabaseConfiguration) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		cfg.Host,
		cfg.User,
		cfg.Password,
		cfg.Name,
		cfg.Port,
		SSL,
		Timezone,
	)
}
