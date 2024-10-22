package model

type FilmRequest struct {
	Title              string  `json:"title" binding:"required,max=255"`
	Description        string  `json:"description" binding:"required,min=10,max=255"`
	ReleaseYear        uint16  `json:"releaseYear" binding:"required,min=0"`
	LanguageID         int64   `json:"languageId" binding:"required,min=0"`
	OriginalLanguageID int64   `json:"originalLanguageId" binding:"required"`
	RentalDuration     int64   `json:"rentalDuration" binding:"required,min=1"`
	RentalRate         float64 `json:"rentalRate" binding:"required,min=0"`
	Length             int64   `json:"length" binding:"required,min=1"`
	ReplacementCost    float64 `json:"replacementCost" binding:"required,min=0"`
	Rating             string  `json:"rating" binding:"required,rating"`
	SpecialFeatures    string  `json:"specialFeatures" binding:"required"`
}
