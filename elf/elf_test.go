package elf

import (
	"fmt"
	"testing"
)

func TestArgon(t *testing.T) {
	key, e := derive_key("shaula00", nil)
	if e != nil {
		t.Error(e)
	}
	fmt.Println(key)
}
