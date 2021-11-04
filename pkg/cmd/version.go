package cmd

import (
	"github.com/wanghexie/ipgw"
	"github.com/wanghexie/ipgw/pkg/console"
	"github.com/urfave/cli/v2"
)

var (
	VersionCommand = &cli.Command{
		Name:  "version",
		Usage: "show version and build info",
		Action: func(ctx *cli.Context) error {
			console.InfoF("ipgw %s+%s\n", ipgw.Version, ipgw.Build)
			return nil
		},
	}
)
