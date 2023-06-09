package models

type UserModel struct {
	Id        string `gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()"`
	Username  string `gorm:"column:usernames;index;unique;required"`
	Age       int    `gorm:"column:age"`
	Email     string `gorm:"column:email;index;unique;required;not null"`
	Password  string `gorm:"column:password;required"`
	FieldTest string `gorm:"column:ftest;required"`
}

func (c *UserModel) TableName() string {
	return "users"
}
