package models

type NotEnoughMoneyError struct {
	Message string
}

func (e *NotEnoughMoneyError) Error() string {
	return e.Message
}
