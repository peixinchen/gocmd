package main

import (
    "github.com/peixinchen/gocmd/base"

    // 这里，我想执行 echo.init 但实际用不到 echo，所以怎么弄？
    //"github.com/peixinchen/gocmd/commands/echo"
)

func main() {
    base.ARunner.List()
}
