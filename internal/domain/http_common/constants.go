package httpcommon

type errorResponseCode struct {
	InvalidRequest      string
	InternalServerError string
	RecordNotFound      string
	MissingIdParameter  string
	InvalidDataType     string
}

var ErrorResponseCode = errorResponseCode{
	InvalidRequest:      "INVALID_REQUEST",
	InternalServerError: "INTERNAL_SERVER_ERROR",
	RecordNotFound:      "RECORD_NOT_FOUND",
	MissingIdParameter:  "MISSING_ID_PARAMETER",
	InvalidDataType:     "INVALID_DATA_TYPE",
}

type customValidationErrCode map[string]string

var CustomValidationErrCode = customValidationErrCode{
	"filmrequest.rating":          "INVALID_FILM_RATING",
	"filmrequest.specialfeatures": "INVALID_SPEACIAL_FEATURES",
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
