package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	_ "github.com/chapgx/elf/cmd"
	"github.com/chapgx/elf/elf"
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
		initstate()
		user := elf.GetUser()
		for {
			reader := bufio.NewReader(os.Stdin)
			fmt.Printf("(%s) %s ", user.Username, "elf❯")
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

func initstate() {
	root := rhombi.Root()
	e := elf.EnvState()
	if e == elf.ErrRootIsNotComplete {
		fmt.Println("The root profile is not completed please complete it before you continue.")
		fmt.Println("Your now root")
	outer:
		for _, sub := range root.Subs {
			if sub.Name == "create" {
				for _, innersub := range sub.Subs {
					if innersub.Name == "masterkey" {
						e := innersub.Run()
						if e != nil {
							panic(e)
						}
						break outer
					}
				}
			}
		}
		e = nil
	}

	if e != nil {
		panic(e)
	}
}
