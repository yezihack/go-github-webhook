package main

import "github.com/urfave/cli/v2"

func Commands() []*cli.Command {
	data := []*cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "init config",
		},
	}
	return data
}
func Flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{Name: "path", Aliases: []string{"u"}, Value: "hook", Usage: "url path, eg: protocal://hostname[:port]/path/"},
		&cli.StringFlag{Name: "bash", Aliases: []string{"b"}, Value: "", Usage: "Execute the script path. eg: /home/hook.sh"},
		&cli.IntFlag{Name: "port", Aliases: []string{"p"}, Value: 6666, Usage: "http port"},
		&cli.StringFlag{Name: "secret", Aliases: []string{"s"}, Value: "", Usage: "github hook secret"},
		&cli.BoolFlag{Name: "debug", Aliases: []string{"bug"}, Value: false, Usage: "print debug info"},
	}
}
