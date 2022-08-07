package customerror

import "fmt"

// ValidationError represent custom error that related to bad request
type (
	ValidationError string
	ErrorNotFound   string
)

func (e ValidationError) Error() string {
	return string(e)
}

func ValidationErrorf(format string, a ...interface{}) ValidationError {
	return ValidationError(fmt.Sprintf(format, a...))
}

func (e ErrorNotFound) Error() string {
	return string(e)
}

func ErrorNotFoundf(format string, a ...interface{}) ErrorNotFound {
	return ErrorNotFound(fmt.Sprintf(format, a...))
}
