package elf

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/chapgx/elf/db"
	"golang.org/x/crypto/argon2"
)

type Password struct {
	cleartext   string
	algo        string
	version     int
	iterations  int
	memory      int
	parallelism int
	salt        []byte
	key         []byte
	hash        string
	redacted    bool // is cleartext redacted
}

// Redacts cleartext passoword for security
func (p *Password) redact() {
	p.cleartext = "*"
	for range 10 {
		p.cleartext += "*"
	}
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

	str += fmt.Sprintf("%x", psswd.salt)

	return str
}

func (pass *Password) Hash() string {
	pass.hash = pass.String()
	return pass.hash
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

// Store saves the password used for the derive key in a hashed form
func (psswd *Password) Store(user string) error {
	fingerp, e := psswd.fingerprent()
	if e != nil {
		return e
	}

	client := db.Connect(_dbpath)
	defer client.Close()

	_, e = client.Exec(`
		update admins
		set fingerprint = ?
		where fingerprint is null
		and uname = ?;
		`, fingerp, user)

	if e == nil {
		psswd.redact()
	}

	return e
}

func (pss *Password) fingerprent() (string, error) {
	if pss.hash == "" {
		return "", errors.New("master key needs bot hash before fingreprint is created")
	}

	sha := sha256.New()
	sha.Write([]byte(pss.hash))
	sha.Write([]byte(pss.cleartext))
	hashed := sha.Sum(nil)

	strhash := fmt.Sprintf("sha256$%x", hashed)

	return strhash, nil
}

func (pss *Password) parse_cleartext_hash(hash string) []byte {
	return nil
}

func NewPassword(hash, cleartext string) Password {
	return Password{algo: "argon2i", hash: hash, cleartext: cleartext}
}
