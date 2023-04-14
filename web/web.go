package web

import (
	"fmt"

	"github.com/leon19951027/woofwoofgpt/config"
)

type Web struct {
	Host string
	Port int
}

func (w *Web) ApplyCfg(cfg *config.Cfg) {
	w.Host = cfg.Host
	w.Port = cfg.Port
	fmt.Println(w)
}

func (w *Web) Run() {

}
