package entity

import "time"

type RefreshToken struct {
	Id         int64     `gorm:"column:id;primaryKey" json:"id"`
	Token      string    `gorm:"column:token" json:"token"`
	Username   string     `gorm:"column:username" json:"username"`
	ExpTime    time.Time `gorm:"column:exp_time;autoUpdateTime" json:"expTime"`
}

func (RefreshToken) TableName() string {
	return "refresh_token"
}
