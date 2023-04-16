package config

import (
	"github.com/spf13/viper"
)

type Server struct {
	Host string
	Port string
}

type GlobalOpenai struct {
	UrlPrefix string
	Token     string
}

type Cfg struct {
	Server
	GlobalOpenai
}

func ReadConf() *Cfg {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	cfg := &Cfg{}
	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		//log.Fatalf("Error reading config file, %s", err)
		panic(err)
	}
	cfg.Server.Host = viper.GetString("server.host")
	cfg.Server.Port = viper.GetString("server.port")
	cfg.GlobalOpenai.UrlPrefix = viper.GetString("globalOpenai.urlPrefix")
	return cfg
}
