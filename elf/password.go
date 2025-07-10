package elf

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/crypto/argon2"
)

type Password struct {
	algo        string
	version     int
	iterations  int
	memory      int
	parallelism int
	salt        []byte
	key         []byte
	hash        string
}

func (psswd Password) String() string {
	str := psswd.algo + "$"
	if psswd.version == 0 {
		panic("version is required unable to parse to string")
	}

	str += fmt.Sprintf("%d$", psswd.version)

	if psswd.iterations == 0 {
		panic("iterations can't be 0 unable to parse to string")
	}

	str += fmt.Sprintf("%d$", psswd.iterations)

	if psswd.memory == 0 {
		panic("memory can't be 0 unable to parse to string")
	}

	str += fmt.Sprintf("%d$", psswd.memory)

	if psswd.parallelism == 0 {
		panic("parallelism can't be 0 unable to parse to string")
	}

	str += fmt.Sprintf("%d$", psswd.parallelism)

	if psswd.salt == nil {
		panic("password salt is <nil>")
	}

	str += fmt.Sprintf("%x$", psswd.salt)

	return str
}

func (psswd *Password) parse_hash() error {
	if psswd.hash == "" {
		return errors.New("hash_str is required")
	}

	parts := strings.Split(psswd.hash, "$")

	if len(parts) < 6 {
		return ErrMalformedPsswd
	}

	version, e := strconv.Atoi(parts[VerIdx])
	if e != nil {
		return e
	}

	if version != argon2.Version {
		return ErrWrongAlgoVersion
	}
	psswd.version = version

	iters, e := strconv.Atoi(parts[IterIdx])
	if e != nil {
		return e
	}
	psswd.iterations = iters

	mem, e := strconv.Atoi(parts[MemIdx])
	if e != nil {
		return e
	}
	psswd.memory = mem

	parall, e := strconv.Atoi(parts[ParallIdx])
	if e != nil {
		return e
	}
	psswd.parallelism = parall

	salt, e := hex.DecodeString(parts[SaltIdx])
	if e != nil {
		return e
	}
	psswd.salt = salt

	return nil
}

func NewPassword(hash string) Password {
	p := Password{algo: "argon2i"}
	if hash != "" {
		p.hash = hash
	}
	return p
}
