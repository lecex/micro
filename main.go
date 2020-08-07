package main

import (
	"github.com/micro/micro/v2/cmd"
	_ "github.com/micro/go-plugins/registry/nats"
)

func main() {
	cmd.Init()
}
