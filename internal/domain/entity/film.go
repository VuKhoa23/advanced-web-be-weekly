package entity

import "time"

type Film struct {
	ID                 int64     `gorm:"column:film_id;primaryKey" json:"id"`
	Title              string    `gorm:"column:title" json:"title"`
	Description        string    `gorm:"column:description" json:"description"`
	ReleaseYear        uint16    `gorm:"column:release_year" json:"releaseYear"`
	LanguageID         int64     `gorm:"column:language_id" json:"languageId"`
	OriginalLanguageID int64     `gorm:"column:original_language_id" json:"originalLanguageId"`
	RentalDuration     int64     `gorm:"column:rental_duration" json:"rentalDuration"`
	RentalRate         float64   `gorm:"column:rental_rate" json:"rentalRate"`
	Length             int64     `gorm:"column:length" json:"length"`
	ReplacementCost    float64   `gorm:"column:replacement_cost" json:"replacementCost"`
	Rating             string    `gorm:"column:rating" json:"rating"`
	SpecialFeatures    string    `gorm:"column:special_features" json:"specialFeatures"`
	LastUpdate         time.Time `gorm:"column:last_update;autoUpdateTime" json:"lastUpdate"`
}

func (Film) TableName() string {
	return "film"
}
