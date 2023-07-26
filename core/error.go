package core

type UserError struct {
	message  string
	original error
}

func NewUserError(err error) *UserError {
	return &UserError{
		message:  err.Error(),
		original: err,
	}
}

func (err *UserError) Error() string {
	return err.message
}

func (err *UserError) Cause() error {
	return err.original
}
