package entity

type User struct {
	Username string `gorm:"column:username;primaryKey" json:"username"`
	Password string `gorm:"column:password" json:"password"`
}

func (User) TableName() string {
	return "user"
}
