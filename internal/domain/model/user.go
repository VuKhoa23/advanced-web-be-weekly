package model

type UserRequest struct {
	UserName string `json:"userName" binding:"required,min=1,max=255"`
	Password string `json:"password" binding:"required,min=1,max=255"`
}
