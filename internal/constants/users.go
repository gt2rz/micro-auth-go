package constants

import "errors"

// errUserNotSaved is the error returned when a user is not saved
var ErrUserNotSaved = errors.New("User not saved")

// errUserNotFound is the error returned when a user is not found
var ErrUserNotFound = errors.New("User not found")

// errGenerateRandomString is the error returned when a random string cannot be generated
var ErrGenerateRandomString = errors.New("Error generating random string")

// errResetTokenNotSaved is the error returned when a reset token is not saved
var ErrResetTokenNotSaved = errors.New("Reset token not saved")

// errUserAlreadyExists is the error returned when a user already exists
var ErrUserAlreadyExists = errors.New("User already exists")