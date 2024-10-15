package model

type FilmRequest struct {
	Title              string  `json:"title" binding:"required"`
	Description        string  `json:"description" binding:"required"`
	ReleaseYear        uint16  `json:"releaseYear" binding:"required"`
	LanguageID         int64   `json:"languageId" binding:"required"`
	OriginalLanguageID int64   `json:"originalLanguageId" binding:"required"`
	RentalDuration     int64   `json:"rentalDuration" binding:"required"`
	RentalRate         float64 `json:"rentalRate" binding:"required"`
	Length             int64   `json:"length" binding:"required"`
	ReplacementCost    float64 `json:"replacementCost" binding:"required"`
	Rating             string  `json:"rating" binding:"required"`
	SpecialFeatures    string  `json:"specialFeatures" binding:"required"`
}