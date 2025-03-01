package error

import "errors"

var (
	ErrorUserNotFound         = errors.New("user not found")
	ErrorPasswordIncorrect    = errors.New("password is incorrect")
	ErrorUsernameExist        = errors.New("username already exist")
	ErrorPasswordDoesNotMatch = errors.New("password does not match")
)

var UserErrors = []error{
	ErrorUserNotFound,
	ErrorPasswordIncorrect,
	ErrorUsernameExist,
	ErrorPasswordDoesNotMatch,
}
