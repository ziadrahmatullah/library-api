package migration

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	a := &entity.Author{}
	b := &entity.Book{}

	_ = db.Migrator().DropTable(a)
	_ = db.Migrator().DropTable(b)

	_ = db.AutoMigrate(a, b)
}
