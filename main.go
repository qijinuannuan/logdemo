package main

import (
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"logtest/g"
	"logtest/hello"
	"os"
)

func main() {
	cfg := flag.String("c", "cfg.json", "config file")
	version := flag.Bool("v", false, "show version")
	flag.Parse()

	if *version {
		fmt.Println("version: No.1")
		os.Exit(0)
	}

	g.ParseConfig(*cfg)
	g.InitLog(g.Config().LogLevel)
	hello.Hello()
	log.Info("start ... ")
}
