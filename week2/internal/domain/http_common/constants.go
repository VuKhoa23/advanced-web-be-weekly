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
