package entity

import "time"

type RefreshToken struct {
	Token    string    `gorm:"column:token" json:"token"`
	Username string    `gorm:"column:username;primaryKey" json:"username"`
	ExpTime  time.Time `gorm:"column:exp_time" json:"expTime"`
}

func (RefreshToken) TableName() string {
	return "refresh_token"
}
