package web

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/leon19951027/woofwoofgpt/config"
	"github.com/leon19951027/woofwoofgpt/web/controller"
)

type Web struct {
	Host string
	Port string
}

func (w *Web) ApplyCfg(cfg *config.Cfg) {
	w.Host = cfg.Host
	w.Port = cfg.Port
	fmt.Println(w)
}

func (w *Web) Run() {
	h := gin.Default()
	v1 := h.Group("/api/v1")
	v1.POST("/strem-chat", controller.Strem_Chat)
	h.Run(w.Host + ":" + w.Port)
}
