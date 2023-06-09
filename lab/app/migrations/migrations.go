package migrations

import (
	"lab/app/models"
	"log"

	"gorm.io/gorm"
)

func Migrator(db *gorm.DB) error {
	// Auto create table if object is't exist
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

	// cl, _ := db.Migrator().ColumnTypes(&user)

	// for _, c := range cl {
	// 	if c.Name() == "ftest" {
	// 		mg := db.Migrator()
	// 		sf := &schema.Field{
	// 			Name:     "ftests",
	// 			DataType: "string",
	// 			Size:     60,
	// 			// DBName:   db.Migrator().CurrentDatabase(),
	// 			DBName:            "ftest",
	// 			Schema:            db.Statement.Schema,
	// 			IndirectFieldType: reflect.TypeOf(user.FieldTest),
	// 		}
	// 		err := mg.MigrateColumn(&user, sf, c)
	// 		if err != nil {
	// 			log.Fatal(err)
	// 			return err
	// 		}
	// 	}
	// }

	// migrator.Migrate(&models.UserModel{}, func(tx *gorm.DB) error {
	// 	// Menjalankan pernyataan SQL untuk mengubah tipe data kolom
	// 	if err := tx.Exec("ALTER TABLE users ALTER COLUMN column_name TYPE new_data_type").Error; err != nil {
	// 		return err
	// 	}

	// 	return nil
	// })
	return nil
}
