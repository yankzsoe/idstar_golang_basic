package crud

import (
	"lab/app/config"
	"lab/app/models"
)

func UpdateSample() {
	db := config.GetDB()
	user := models.UserModel{}

	db.Model(&user).Where("age = ?", 27).Update("usernames", "Jhon")
}

func UpdateSample2() {
	db := config.GetDB()

	data := models.UserModel{
		Username: "Luge",
		Email:    "test@email.com",
	}

	if err := db.Model(&models.UserModel{}).Where("age = ?", 21).Updates(data).Error; err != nil {
		panic(err)
	}
}
