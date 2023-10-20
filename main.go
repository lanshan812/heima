package main

import (
	"log"
	"net/http"
	"patent/bootstrap"
	routers "patent/router"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// @host 120.92.116.143
// @BasePath /api
// doc url: http://120.92.116.143：5151/swagger/index.html#/
func main() {
	//初始化配置文件，log, 数据库，redis
	bootstrap.Init()

	//创建gin实例
	gin.SetMode(viper.GetString("RunMode"))
	g := gin.New()
	routers.Load(g)

	//启动服务
	port := viper.GetString("ServicePort")
	log.Println("patent service start, listening requests on http address: " + port)
	log.Println(http.ListenAndServe(port, g).Error())

}
