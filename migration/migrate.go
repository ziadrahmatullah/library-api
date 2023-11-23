package migration

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	b := &entity.Book{}

	_ = db.Migrator().DropTable(b)

	_ = db.AutoMigrate(b)
}
