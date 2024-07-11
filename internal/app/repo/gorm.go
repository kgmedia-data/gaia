package repo

import (
	"database/sql"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GormRepo struct {
	GormDB *gorm.DB
	Client *sql.DB
}

func NewGormRepo(path string, nConn int) (*GormRepo, error) {
	var err error
	database := new(GormRepo)
	if database.Client, err = sql.Open("postgres", path); err != nil {
		return nil, err
	}

	database.GormDB, err = gorm.Open(postgres.New(postgres.Config{Conn: database.Client}),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Warn),
		})
	if err != nil {
		return nil, err
	}

	db, err := database.GormDB.DB()
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(nConn)
	db.SetMaxOpenConns(nConn)
	db.SetConnMaxLifetime(1 * time.Hour)

	return database, nil
}
