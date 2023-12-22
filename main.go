package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"model-k8s-service-go/api"
	"model-k8s-service-go/config"
	"model-k8s-service-go/initialize"
	"model-k8s-service-go/util"
	"net/http"
)

func Hello(c *gin.Context) {
	c.String(200, "hello %s", "world")
	c.JSON(http.StatusOK, gin.H{ //以json格式输出
		"name": "tom",
		"age":  "20",
	})
}

func main() {
	initialize.Init("config.yaml")
	r := gin.Default() //创建一个默认的路由引擎
	r.Use(gin.CustomRecovery(util.Recovery))
	api.Api(r)
	fmt.Println(config.Config.AppPort)
	r.Run(":" + config.Config.AppPort)
}
