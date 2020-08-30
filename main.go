package main

import (
	_ "github.com/lecex/core/plugins" // 插件在后面执行
	"github.com/micro/micro/v2/cmd"
)

func main() {
	cmd.Init()
}
