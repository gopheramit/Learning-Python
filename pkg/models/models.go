package models

import (
	"errors"
)

var ErrNoRecord = errors.New("models: no matching record found")

type PythonUser struct {
	UserID string
	Email  string
	TaskID int
}

type Tasks struct {
	TaskID          int
	TaskName        string
	TaskDescription string
	Difficulty      string
	Output          string
}
