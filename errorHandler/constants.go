package errorHandler

const (
	InternalErrorDatabase            = "database error"
	InternalErrorMarshaling          = "error during data marshalling"
	InternalErrorNoEmployeeFound     = "no documents found in employee collection"
	InternalErrorInvalidPayload      = "invalid payload calling put employees"
	InternalErrorEmployeeIdNotGiven  = "employee id is missing"
	InternalErrorEmployeeIdNotUnique = "at least one employee id is not unique"
)
const (
	ExternalErrorEmployeeNotFound    = "employee not found"
	ExternalErrorInvalidPayload      = "invalid payload"
	ExternalErrorIdMissingInURL      = "id is missing in URL"
	ExternalErrorEmployeeIdNotUnique = "at least one employee id is not unique"
	ExternalErrorUnknown             = "unknown error"
)
