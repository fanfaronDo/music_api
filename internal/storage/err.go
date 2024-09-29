package storage

import "errors"

var (
	ErrConfigNull         = errors.New("config value should not be null")
	ErrConfigPortLessZero = errors.New("config port should greater than zero")
)
