package main

import (
	"github.com/wanghexie/ipgw/pkg/cmd"
	"github.com/wanghexie/ipgw/pkg/console"
	"os"
)

func main() {
	if err := cmd.App.Run(os.Args); err != nil {
		console.FatalL(err.Error())
	}
}
