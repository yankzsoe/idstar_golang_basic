package migrations

import (
	"lab/app/models"
	"log"

	"gorm.io/gorm"
)

func Migrator(db *gorm.DB) error {
	// Auto create table if user object is't exist
	user := models.UserModel{}
	if !db.Migrator().HasTable(user.TableName()) {
		db.AutoMigrate(user)
	}

	// Rename column
	if !db.Migrator().HasColumn(user, "ftests") {
		if err := db.Migrator().RenameColumn(&models.UserModel{}, "ftest", "ftests"); err != nil {
			log.Fatal(err)
			return err
		}
	}

	return nil
}
