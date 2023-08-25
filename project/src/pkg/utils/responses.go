package utils

const Wrong_password = "WRONG_PASSWORD"
const Not_found = "NOT_FOUND"
const Server_Error = "SERVER_ERROR"

type WrongPassword struct{}
type NotFound struct{}
type ServerError struct {
	Detail *string
}

func (e *ServerError) Error() string {
	return Server_Error + " : " + *e.Detail
}
func (e *NotFound) Error() string {
	return Not_found
}

func (e *WrongPassword) Error() string {
	return Wrong_password
}

func NewWrongPasswordError() error {
	return &WrongPassword{}
}
func NewNotFoundError() error {
	return &NotFound{}
}

func NewServerError(detail *string) error {
	return &ServerError{Detail: detail}
}
