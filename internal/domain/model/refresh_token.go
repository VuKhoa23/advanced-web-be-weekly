package model

import "time"

type RefreshTokenRequest struct {
	Token   string    `json:"token" binding:"required"`
	Username  string     `json:"username" binding:"required"`
	ExpTime time.Time `json:"expTime" binding:"required"`
}