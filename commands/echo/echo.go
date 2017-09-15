package echo

import (
    "fmt"

    "github.com/peixinchen/gocmd/base"
)

type EchoCommand struct {
    base.BaseCommand
}

func (c EchoCommand) Execute() {
    fmt.Printf("{%v}: Execute\n", c.Name)
}

func init() {
    base.RegisterCommand(EchoCommand{
        base.BaseCommand{
            Name: "echo",
            Short:  "回声",
            Long:   "你说什么，我就回什么。这就是回声。",
        },
    })
}
