package repo

import (
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

var (
	TEST_CONN_STRING = "postgres://gaia:gaia123@localhost:5432/gaia?sslmode=disable"
)

// function to truncate tables with the parameter being GormDB, array of table names
func truncateTables(db *gorm.DB, tables ...string) {
	for _, table := range tables {
		err := db.Exec("TRUNCATE TABLE " + table + " RESTART IDENTITY CASCADE;").Error
		if err != nil {
			log.Info("No error")
		}
	}
}
