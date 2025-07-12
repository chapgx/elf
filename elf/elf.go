/**
* package elf packages handles core functionality of the elf library
**/
package elf

import (
	"crypto/rand"
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/chapgx/elf/db"
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

var _elfdir = filepath.Join(home(), ".elf")

var _dbpath = filepath.Join(_elfdir, "elf.db")

var _key []byte

var _currentuser = &User{}

func auth(password, hash string) {
}

// Derive encryption key
func DeriveKey(pass *Password, salt []byte) ([]byte, error) {
	return derive_key(pass, salt)
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

	_key = key

	return key, nil
}

// GetDbPath returns a path to the sqlite db

func GetDbPath() string {
	return _dbpath
}

// Get user home directory
func home() string {
	var home string
	switch runtime.GOOS {
	case "darwin", "linux":
		home = os.Getenv("HOME")
	case "windows":
		home = os.Getenv("PROFILE")
	}
	return home
}

// Initializes elf environment
func Init() error {
	if !strings.HasSuffix(_elfdir, ".elf") {
		return errors.New("wrong path to perform action")
	}
	e := os.Mkdir(_elfdir, 0700)

	if e != nil {
		return e
	}

	e = db.Init(_dbpath)

	if e != nil {
		return e
	}

	e = Admin{}.init()

	return e
}

// EnvState reads the current state of the elf environment
func EnvState() error {
	_, e := os.Stat(_elfdir)

	if os.IsNotExist(e) {
		e = Init()
		if e != nil {
			return e
		}
	}

	admin, e := Admin{}.ReadRoot()
	if e != nil {
		return e
	}

	e = admin.IsRootComplete()
	if e == ErrRootIsNotComplete {
		_currentuser.Username = "root"
	}

	return e
}

// Torch Destroys elf environment
func Torch() error {
	if !strings.HasSuffix(_elfdir, ".elf") {
		return errors.New("wrong path to perform action")
	}
	e := os.RemoveAll(_elfdir)
	return e
}

func GetUser() *User {
	return _currentuser
}
