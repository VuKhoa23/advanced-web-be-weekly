package httpcommon

type errorResponseCode struct {
	InvalidRequest      string
	InternalServerError string
	RecordNotFound      string
	MissingIdParameter  string
}

var ErrorResponseCode = errorResponseCode{
	InvalidRequest:      "INVALID_REQUEST",
	InternalServerError: "INTERNAL_SERVER_ERROR",
	RecordNotFound:      "RECORD_NOT_FOUND",
	MissingIdParameter:  "MISSING_ID_PARAMETER",
}

type customValidationErrCode map[string]string

var CustomValidationErrCode = customValidationErrCode{
	"filmrequest.rating": "INVALID_FILM_RATING",
}

type errorMessage struct {
	GormRecordNotFound string
	InvalidDataType    string
	InvalidRequest     string
}

var ErrorMessage = errorMessage{
	GormRecordNotFound: "record not found",
	InvalidDataType:    "invalid data type",
	InvalidRequest:     "invalid request",
}
