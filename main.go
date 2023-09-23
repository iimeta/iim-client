package main

import (
	_ "github.com/iimeta/iim-client/internal/core"

	_ "github.com/iimeta/iim-client/internal/packed"

	_ "github.com/iimeta/iim-client/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/iimeta/iim-client/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
