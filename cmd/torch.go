package cmd

import (
	"fmt"
	"strings"

	"github.com/chapgx/elf/elf"
	"github.com/racg0092/rhombifer"
)

var torchcmd = &rhombifer.Command{
	Name:      "torch",
	ShortDesc: "Cleans up the environment this action if finite and not reversible",
	Run: func(args ...string) error {
		fmt.Print(" Your about to destroy your local environmnet are you sure you want to proceed ? [y/n] ")
		var confirmation string
		_, e := fmt.Scanln(&confirmation)
		if e != nil {
			return e
		}

		confirmation = strings.ToLower(confirmation)
		confirmation = strings.ReplaceAll(confirmation, "\n", "")

		if confirmation == "n" || confirmation == "no" {
			fmt.Println("aborting action")
			return nil
		}

		if confirmation != "y" && confirmation != "yes" {
			fmt.Printf("imput was not recognized <%s>", confirmation)
			return nil
		}

		return elf.Torch()
	},
}

func init() {
	root := rhombifer.Root()

	root.AddSubs(torchcmd)
}
