package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	_ "github.com/chapgx/elf/cmd"
	rhombi "github.com/racg0092/rhombifer"
	"github.com/racg0092/rhombifer/pkg/builtin"
	"github.com/racg0092/rhombifer/pkg/text"
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

	root.Run = func(args ...string) error {
		for {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("elf❯ ")
			input, e := reader.ReadString('\n')
			if e != nil {
				txt := text.SetForGroundColor(text.RED, "󰨰 = "+e.Error())
				fmt.Printf("%s\n", txt)
				continue
			}

			input = strings.ToLower(input)
			input = strings.ReplaceAll(input, "\n", "")
			os.Args = []string{"elfd"}
			os.Args = append(os.Args, strings.Split(input, " ")...)

			e = rhombi.Start()
			if e != nil {
				txt := text.SetForGroundColor(text.RED, "󰨰 = "+e.Error())
				fmt.Printf("%s\n\n", txt)
			}

		}
	}
}
