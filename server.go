package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	_ "github.com/houseme/goframe-polaris-demo/internal/logic"
	_ "github.com/houseme/goframe-polaris-demo/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/houseme/goframe-polaris-demo/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
