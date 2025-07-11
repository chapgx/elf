package elf

import (
	"testing"

	. "github.com/chapgx/assert"
)

func TestArgon(t *testing.T) {
	t.Run("test hashing and parsing", func(t *testing.T) {
		pass := NewPassword("", "shaula00")
		_, e := derive_key(&pass, nil)
		AssertT(t, e == nil, "error was not <nil>: "+e.Error())

		hash := pass.Hash()

		pass2 := NewPassword(hash, "")
		pass2.parse_hash()

		hash2 := pass2.Hash()
		AssertT(t, hash == hash2, "hash should be equal to hash2")
	})
}
