package models

import (
	"errors"
)

var ErrNoRecord = errors.New("models: no matching record found")

type PythonUser struct {
	userID string
	email  string
	TaskID int
}
