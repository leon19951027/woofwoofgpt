package web

import (
	"myopenai/config"
)

type Web struct {
	Host string
	Port int
}

func (w *Web) ApplyCfg(cfg *config.Cfg) {
	w.Host = cfg.Host
	w.Port = cfg.Port
}

func (w *Web) Run() {

}
