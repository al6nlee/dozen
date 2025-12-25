package initialize

import (
	"dozen/backend/internal/global"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitConfig() {
	configFileName := "config/config.yaml"
	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(global.ServerConfig); err != nil {
		panic(err)
	}
	fmt.Printf("配置加载完成, %+v\n", global.ServerConfig)

	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置文件修改了...")
		_ = v.Unmarshal(global.ServerConfig)
		fmt.Printf("配置更新完成, %+v\n", global.ServerConfig)
	})
}
