package model

type FilmRequest struct {
	Title string `json:"title" binding:"required"`
}
