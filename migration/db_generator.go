package migration

import (
	"github.com/Screen17/catalog/appcontext"
	"github.com/Screen17/catalog/model"
)

// CreateDatabase creates the tables used in this application.
func CreateDatabase(context appcontext.Context) {
	if context.GetConfig().Database.Migration {
		db := context.GetRepository()

		db.DropTableIfExists(&model.Account{})
		db.DropTableIfExists(&model.Authority{})

		db.AutoMigrate(&model.Account{})
		db.AutoMigrate(&model.Authority{})
		db.AutoMigrate(&model.Dataset{})
	}
}
