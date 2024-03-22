package helper

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

// 加载配置文件
func LoadConfig() {
	viper.SetConfigName("config")      // 配置文件的名称 (不需要扩展名)
	viper.SetConfigType("toml")        // 设置配置文件类型
	viper.AddConfigPath("data/config") // 添加查找配置文件的路径

	// 你也可以添加更多的路径，例如配置文件在另一个目录
	// viper.AddConfigPath("$HOME/.appname")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
		os.Exit(1)
	}
}
