package err

import "net/http"

var (
	DatabaseError = &Error{HttpStatus: http.StatusInternalServerError, Code: "InternalError", Message: "Internal server error!"}
	Unauthorized  = &Error{HttpStatus: http.StatusUnauthorized, Code: "Unauthorized", Message: "Unauthorized, you need login!"}
	InvalidParam  = &Error{HttpStatus: http.StatusBadRequest, Code: "InvalidParam", Message: "Bad request, invalid param!"}
	NotFound      = &Error{HttpStatus: http.StatusNotFound, Code: "NotFound", Message: "Resource not found!"}
)

// Error defines a standard application error.
type Error struct {
	HttpStatus int
	// Machine-readable error code.
	Code string
	// Human-readable message.
	Message string
	// Logical operation and nested error.
	Action string
	Err    error
}

func New(status int, code, message, action string, err error) *Error {
	return &Error{
		HttpStatus: status,
		Code:       code,
		Message:    message,
		Action:     action,
		Err:        err,
	}
}
