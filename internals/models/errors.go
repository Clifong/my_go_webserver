package models

import (
	"errors"
)

var (
	ErrNoRecord = errors.New("modes: No matching records found")

	ErrInvalidCredentials = errors.New("models: invalid credentials")

	ErrDuplicateEmail = errors.New("models: Duplicate email")
)
