package model

type ActorRequest struct {
	FirstName string `json:"firstName" binding:"required,min=1,max=255"`
	LastName  string `json:"lastName" binding:"required,min=1,max=255"`
}
