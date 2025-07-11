package elf

import (
	"crypto/rand"

	"golang.org/x/crypto/argon2"
)

const (
	Memory      = 64 * 1024
	Iterations  = 3
	Parallelism = 2
	SaltLength  = 16
	KeyLength   = 32
)

const (
	VerIdx = iota + 1
	IterIdx
	MemIdx
	ParallIdx
	SaltIdx
	HashIdx
)

func auth(password, hash string) {
}

// Hashed password
func derive_key(password *Password, salt []byte) (key []byte, err error) {
	if salt == nil {
		salt = make([]byte, SaltLength)
		_, err = rand.Read(salt)
		if err != nil {
			return nil, err
		}
	}

	key = argon2.Key([]byte(password.cleartext), salt, Iterations, Memory, Parallelism, KeyLength)

	password.version = argon2.Version
	password.iterations = Iterations
	password.memory = Memory
	password.parallelism = Parallelism
	password.salt = salt
	password.key = key

	password.redact()

	return key, nil
}
