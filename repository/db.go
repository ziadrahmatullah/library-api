package repository

import (
	"os"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetConnection() (*gorm.DB, error) {
	dsn := util.GetDsn()
	var l = logger.Default.LogMode(logger.Info)
	if os.Getenv("ENV") == "dev" {
		l = nil
	}
	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: l,
	})
}
