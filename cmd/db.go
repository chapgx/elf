package cmd

import (
	"os"
	"os/exec"

	"github.com/chapgx/elf/elf"
	"github.com/racg0092/rhombifer"
)

var dbcmd = &rhombifer.Command{
	Name:      "db",
	ShortDesc: "SQLITE database",
	Run: func(args ...string) error {
		command := exec.Command("sqlite3", elf.GetDbPath())
		command.Stdin = os.Stdin
		command.Stdout = os.Stdout

		e := command.Run()

		return e
	},
}

func init() {
	root := rhombifer.Root()
	root.AddSub(dbcmd)
}
