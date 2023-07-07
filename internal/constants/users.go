package constants

import "errors"

var ErrUserNotSaved = errors.New("User not saved")

var ErrUserNotFound = errors.New("User not found")

var ErrGenerateRandomString = errors.New("Error generating random string")

var ErrResetTokenNotSaved = errors.New("Reset token not saved")

var ErrUserAlreadyExists = errors.New("User already exists")

var ErrGettingUser = errors.New("Error getting user")

var ErrResetTokenExpired = errors.New("Reset token expired")
