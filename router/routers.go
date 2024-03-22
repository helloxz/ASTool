package router

import (
	"astool/controller"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Start() {
	// 运行模式
	model := viper.GetString("app.model")
	// 运行端口
	port := viper.GetString("app.port")

	// 设置运行模式和端口
	gin.SetMode(model)

	// 启动gin
	r := gin.Default()
	r.GET("/", controller.Test)
	r.Run(":" + port)

}
