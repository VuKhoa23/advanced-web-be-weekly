package constants

type errorMessage struct {
	GormRecordNotFound string
}

var ErrorMessage = errorMessage{
	GormRecordNotFound: "record not found",
}
