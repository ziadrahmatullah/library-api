package migration

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	a := &entity.Author{}
	b := &entity.Book{}
	u := &entity.User{}

	_ = db.Migrator().DropTable(a)
	_ = db.Migrator().DropTable(b)
	_ = db.Migrator().DropTable(u)

	_ = db.AutoMigrate(a, b, u)
}
