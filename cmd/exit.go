package cmd

import (
	"os"

	"github.com/racg0092/rhombifer"
)

var exitcmd = &rhombifer.Command{
	Name:      "exit",
	ShortDesc: "Exits REPL",
	Run: func(args ...string) error {
		os.Exit(0)
		return nil
	},
}

func init() {
	root := rhombifer.Root()
	root.AddSub(exitcmd)
}
