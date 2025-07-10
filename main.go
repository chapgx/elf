package main

import (
	"fmt"

	rhombi "github.com/racg0092/rhombifer"
	"github.com/racg0092/rhombifer/pkg/builtin"
)

func main() {
	if err := rhombi.Start(); err != nil {
		fmt.Println(err)
	}
}

func init() {
	root := rhombi.Root()
	help := builtin.HelpCommand(nil, nil)
	root.AddSub(&help)

	config := rhombi.GetConfig()
	config.RunHelpIfNoInput = true
}
