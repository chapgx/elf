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
func derive_key(password string, salt []byte) (key []byte, err error) {
	if salt == nil {
		salt = make([]byte, SaltLength)
		_, err = rand.Read(salt)
		if err != nil {
			return nil, err
		}
	}

	key = argon2.Key([]byte(password), salt, Iterations, Memory, Parallelism, KeyLength)

	p := NewPassword("")
	p.version = argon2.Version
	p.iterations = Iterations
	p.memory = Memory
	p.parallelism = Parallelism
	p.salt = salt
	p.key = key

	// TODO: store password information

	return key, nil
}
