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

var elfdir = filepath.Join(home(), ".elf")

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
	if !strings.HasSuffix(elfdir, ".elf") {
		return errors.New("wrong path to perform action")
	}
	e := os.Mkdir(elfdir, 0700)

	if e != nil {
		return e
	}

	dbpath := filepath.Join(elfdir, "elf.db")

	database := db.Connect(dbpath)
	defer database.Close()

	_, e = database.Exec("PRAGMA journal_mode = WAL;")
	if e != nil {
		return e
	}

	_, e = database.Exec(`
	create table if not exists admin (
		id INTEGER PRIMARY KEY,
		uname TEXT,
		masterkey TEXT
	)
	`)

	return e
}

// Torch Destroys elf environment
func Torch() error {
	if !strings.HasSuffix(elfdir, ".elf") {
		return errors.New("wrong path to perform action")
	}
	e := os.RemoveAll(elfdir)
	return e
}
