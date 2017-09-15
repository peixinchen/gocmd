package main

import (
    "github.com/peixinchen/gocmd/base"

    _ "github.com/peixinchen/gocmd/commands/echo"
)

func main() {
    base.ARunner.Run("echo")
}
