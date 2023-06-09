package crud

import (
	"fmt"
	"lab/app/config"
	"lab/app/models"
	"lab/common"
	"math/rand"
)

func getDummy() models.UserModel {
	ids := rand.Intn(100)
	user := models.UserModel{
		Username: "Name" + common.IntToStr(ids),
		Email:    "email" + common.IntToStr(ids),
		Password: common.IntToStr(ids),
	}
	return user
}

func CreateRecord() {
	user := getDummy()

	db := config.GetDB()
	result := db.Create(&user)
	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println(user)
}

func CreateWithSelectField() {
	user := getDummy()

	db := config.GetDB()
	db.Select("usernames", "email", "password").Create(&user)
}
