package model

type HttpResponse[T any] struct {
	Message string `json:"message"`
	Data    *T     `json:"data"`
}
