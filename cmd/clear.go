package cmd

import (
	"fmt"

	"github.com/racg0092/rhombifer"
)

var clearcmd = &rhombifer.Command{
	Name:      "clear",
	ShortDesc: "Clears terminal",
	Run: func(args ...string) error {
		fmt.Println("\x1b[2J\x1b[H")
		return nil
	},
}

func init() {
	root := rhombifer.Root()
	root.AddSub(clearcmd)
}
