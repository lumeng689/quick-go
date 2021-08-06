package cfg

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var v *viper.Viper

func SetUp(config string) {
	v = viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
	})
}
