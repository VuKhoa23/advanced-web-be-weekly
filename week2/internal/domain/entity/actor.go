package entity

import "time"

type Actor struct {
	Id         int64     `gorm:"column:actor_id;primaryKey" json:"id"`
	FirstName  string    `gorm:"column:first_name" json:"firstName"`
	LastName   string    `gorm:"column:last_name" json:"lastName"`
	LastUpdate time.Time `gorm:"column:last_update;autoUpdateTime" json:"lastUpdate"`
}

func (Actor) TableName() string {
	return "actor"
}
