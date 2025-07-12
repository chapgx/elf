package elf

import "errors"

var (
	ErrMalformedPsswd    = errors.New("malformed password hash")
	ErrWrongAlgoVersion  = errors.New("wrong argon2 algorithm version")
	ErrEnvNotSetUp       = errors.New("error environment set up not finish")
	ErrRootIsNotComplete = errors.New("root admin is not complete")
)
