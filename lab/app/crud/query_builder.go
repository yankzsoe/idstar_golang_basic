package crud

import (
	"fmt"
	"lab/app/config"
	"lab/app/models"
)

func QueryBuilderRawSample() {
	db := config.GetDB()
	users := []models.UserModel{}

	rows, err := db.Raw("SELECT * FROM users LIMIT 5").Rows()
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		err := db.ScanRows(rows, &users)
		if err != nil {
			panic(err)
		}
	}

	for i := 0; i < len(users); i++ {
		fmt.Println(users[i])
	}
}

func QueryBuilderExecSample() {
	db := config.GetDB()

	result := db.Exec("UPDATE users SET \"password\"=? WHERE \"usernames\"=?", "qwerty123", "Alma")
	if result.Error != nil {
		panic(result.Error)
	}
}
