package main

import (
	"github.com/leon19951027/woofwoofgpt/config"
	"github.com/leon19951027/woofwoofgpt/web"
)

func main() {
	cfg := config.ReadConf()
	websvc := &web.Web{}
	websvc.ApplyCfg(cfg)
	websvc.Run()

}
