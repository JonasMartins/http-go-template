package utils

const Wrong_password = "WRONG_PASSWORD"
const Not_found = "NOT_FOUND"

type WrongPassword struct{}
type NotFound struct{}

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
