package main

import (
	"github.com/gogf/gf/v2/os/gtime"
	_ "github.com/iimeta/iim-client/internal/core"

	_ "github.com/iimeta/iim-client/internal/packed"

	_ "github.com/iimeta/iim-client/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/iimeta/iim-client/internal/cmd"
)

func main() {

	// 设置进程全局时区
	err := gtime.SetTimeZone("Asia/Shanghai")
	if err != nil {
		panic(err)
	}

	cmd.Main.Run(gctx.GetInitCtx())
}
