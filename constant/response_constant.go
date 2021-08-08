package constant

type ResponseStatus string

const (
	// ResponseStatusSuccessText for
	ResponseStatusSuccessText ResponseStatus = "success"
	// ResponseStatusCreatedText for
	ResponseStatusCreatedText ResponseStatus = "success insert data"
	// ResponseStatusConflictText for
	ResponseStatusConflictText ResponseStatus = "conflict"
	// ResponseStatusInternalServerErrorText for
	ResponseStatusInternalServerErrorText ResponseStatus = "internal server error"
	// ResponseStatusBadRequestText for
	ResponseStatusBadRequestText ResponseStatus = "bad request"
	// ResponseStatusNotFoundText for
	ResponseStatusNotFoundText ResponseStatus = "not found"
	// ResponseStatusUnprocessableEntityText for
	ResponseStatusUnprocessableEntityText ResponseStatus = "unprocessable entity"
	// ResponseStatusUnauthorized for
	ResponseStatusUnauthorized ResponseStatus = "unauthorized"
	// ResponseStatusForbidden for
	ResponseStatusForbidden ResponseStatus = "forbidden"
)
