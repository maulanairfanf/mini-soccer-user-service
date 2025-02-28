package error 

import "errors"

var {
	ErrorUserNotFound = errors.New(text: "user not found")
	ErrorPasswordIncorrect  = errors.New(text: "password is incorrect")
	ErrorUsernameExist = errors.New(text: "username already exist")
	ErrorPasswordDoesNotMatch = errors.New(text: "password does not match")
}

var UserErrors = []error {
	ErrorUserNotFound,
	ErrorPasswordIncorrect,
	ErrorUsernameExist,
	ErrorPasswordDoesNotMatch,
}