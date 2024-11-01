package entity

type User struct {
	UserName string `gorm:"column:user_name;primaryKey" json:"userName"`
	Password string `gorm:"column:password" json:"password"`
}

func (User) TableName() string {
	return "user"
}
