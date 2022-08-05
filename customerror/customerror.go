package customerror

import "fmt"

// ValidationError represent custom error that related to bad request
type (
	ValidationError string
)

func (e ValidationError) Error() string {
	return string(e)
}

func ValidationErrorf(format string, a ...interface{}) ValidationError {
	return ValidationError(fmt.Sprintf(format, a...))
}
