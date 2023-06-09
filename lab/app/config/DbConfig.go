package config

import (
	"lab/app/migrations"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var dB *gorm.DB

var customLogger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	logger.Config{
		LogLevel: logger.Info, // Log level
	},
)

func InitDb() {
	// dsn := "host=localhost user=postgres password=Admin123$ dbname=lab port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	dsn := "host=localhost user=postgres password=Admin123$ dbname=lab port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   customLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
	})

	if err != nil {
		log.Fatal("Error when open connection to database: ", err.Error())
		panic(err)
	}

	db.Debug()

	err = migrations.Migrator(db)
	if err != nil {
		panic(err)
	}

	dB = db
}

func GetDB() *gorm.DB {
	return dB
}
