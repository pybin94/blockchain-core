package main

import (
	"block_chain/app"
	"block_chain/config"
	"flag"
)

var (
	confingFlag = flag.String("environment", "./environment.toml", "environment toml file not found")
	difficulty  = flag.Int("difficuty", 12, "difficulty err")
)

func main() {
	flag.Parse()
	c := config.NewConfig(*confingFlag)
	app.NewApp(c)
}
