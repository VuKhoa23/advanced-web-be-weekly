package httpcommon

type HttpResponse[T any] struct {
	Success bool    `json:"success"`
	Data    *T      `json:"data"`
	Errors  []Error `json:"errors"`
}

type Error struct {
	Message string `json:"message"`
	Code    string `json:"code"`
	Field   string `json:"field"`
}

func NewErrorResponse(error ...Error) HttpResponse[any] {
	return HttpResponse[any]{
		Success: false,
		Data:    nil,
		Errors:  error,
	}
}

func NewSuccessResponse[T any](data *T) HttpResponse[T] {
	return HttpResponse[T]{
		Success: true,
		Data:    data,
		Errors:  nil,
	}
}
