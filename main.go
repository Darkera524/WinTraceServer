package main

import (
	"github.com/Darkera524/WinTraceServer/rpc"
	"github.com/Darkera524/WinTraceServer/g"
	"flag"
)

func main() {
	cfg := flag.String("c", "cfg.json", "configuration file")
	g.ParseConfig(*cfg)

	go g.CronParse(*cfg)
	go rpc.Start()

	select {}

}
