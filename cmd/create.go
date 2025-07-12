package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/chapgx/elf/elf"
	"github.com/racg0092/rhombifer"
	"golang.org/x/term"
)

var createcmd = &rhombifer.Command{
	Name:      "create",
	ShortDesc: "Creates a resource",
}

var createsub_masterkey = &rhombifer.Command{
	Name:      "masterkey",
	ShortDesc: "Creates a master encryption key",
	Run: func(args ...string) error {
		fmt.Print("󱕴 Master Key: ")
		pass1, epass1 := term.ReadPassword(int(os.Stdin.Fd()))

		fmt.Print("\n󱕴 Master Key Verification: ")
		pass2, epass2 := term.ReadPassword(int(os.Stdin.Fd()))
		fmt.Println("")

		if e := errors.Join(epass1, epass2); e != nil {
			return e
		}

		if string(pass1) != string(pass2) {
			return errors.New("password do not macth")
		}

		pass := elf.NewPassword("", string(pass1))

		_, e := elf.DeriveKey(&pass, nil)
		if e != nil {
			return e
		}

		e = elf.Admin{}.SetKey(pass.Hash())

		return e
	},
}

func init() {
	root := rhombifer.Root()
	root.AddSub(createcmd)

	createcmd.AddSubs(
		createsub_masterkey,
	)
}
