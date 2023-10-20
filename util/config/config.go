package config

import (
	"log"
	"os"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Name string
}

func Init(cfg string) {
	c := Config{Name: cfg}

	if err := c.initConfig(); err != nil {
		log.Fatal("Init config error:  " + err.Error())
	}
	//监控配置文件变化并热加载程序
	c.watchConfig()
}

func (c *Config) initConfig() error {

	viper.AddConfigPath("./config")
	viper.SetConfigType("yaml")

	//需要设置环境变量export Patent_ENV="production"
	env := os.Getenv("Patent_ENV")
	if env != "" {
		viper.SetConfigName(env)
	} else {
		log.Fatal("Can not read config file, env variable Patent_ENV is not to be set. will use default config file(default.yaml).")
		viper.SetConfigName("dev")
	}

	cfgFile := "./config/" + env + ".yaml"
	// 如果指定了配置文件，则解析指定的配置文件
	if c.Name != "" {
		cfgFile = c.Name
		viper.SetConfigFile(c.Name)
	}

	log.Printf("Env viriable Patent_ENV is %s, config file name is %s", env, cfgFile)

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

// 监控配置文件变化并热加载程序
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s", e.Name)
	})
}
