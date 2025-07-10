package elf

import "errors"

var (
	ErrMalformedPsswd   = errors.New("malformed password hash")
	ErrWrongAlgoVersion = errors.New("wrong argon2 algorithm version")
)
