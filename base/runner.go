package base

import (
    "fmt"
)

type Runner struct {
    commands []Command
}

func (r Runner) query(name string) Command {
    for _, command := range r.commands {
        if command.GetName() == name {
            return command
        }
    }

    panic("没有找到命令")
}

func (r Runner) List() {
    for _, command := range r.commands {
        fmt.Println(command.GetName())
    }
}

func (r Runner) Run(name string) {
    command := r.query(name)
    
    command.PreExecute()
    command.Execute()
    command.PostExecute()
}

var ARunner Runner

func init() {
    ARunner = Runner{}
}

func RegisterCommand(command Command) {
    ARunner.commands = append(ARunner.commands, command)
}
