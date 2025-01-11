package models

import (
	"errors"
)

var ErrNoRows = errors.New("models: No matching records found")
