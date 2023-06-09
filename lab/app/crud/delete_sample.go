package crud

import (
	"fmt"
	"lab/app/config"
	"lab/app/models"

	"gorm.io/gorm/clause"
)

func DeleteSample() {
	db := config.GetDB()

	if err := db.Delete(&models.UserModel{}, "usernames=? AND age=?", "Jhon", 27).Error; err != nil {
		panic(err)
	}
}

func DeleteSample2() {
	db := config.GetDB()
	user := models.UserModel{}

	if err := db.Clauses(clause.Returning{}).Delete(&user, "usernames=? AND email=?", "Alma", "email@test.com").Error; err != nil {
		panic(err)
	}

	fmt.Println("Data yang dihapus: ", user)
}
