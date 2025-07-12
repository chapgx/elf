package cmd

import (
	"github.com/chapgx/elf/elf"
	"github.com/racg0092/rhombifer"
)

var initcommand = &rhombifer.Command{
	Name:      "init",
	ShortDesc: "initialized elf environment",
	Run: func(args ...string) error {
		e := elf.Init()
		return e
	},
}

func init() {
	root := rhombifer.Root()

	root.AddSubs(initcommand)
}
