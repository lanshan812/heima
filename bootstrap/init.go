package bootstrap

import (
	"patent/model"
	"patent/util/config"
	"patent/util/log"

	"github.com/spf13/pflag"
)

func Init() {
	cfg := pflag.StringP("config", "c", "", "set patent config file path")

	pflag.Parse()
	initConfig(*cfg)
	initLog()
	initDB()
}
func initConfig(cfg string) {
	config.Init(cfg)
}
func initLog() {
	log.InitLog()
}

func initDB() {
	model.InitDB()
}
