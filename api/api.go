package api

import (
	"github.com/gin-gonic/gin"
	"model-k8s-service-go/controller"
	"model-k8s-service-go/util"
)

func Api(e *gin.Engine) {
	trainGroup := e.Group("/api/train")
	trainGroup.POST("", util.Run(controller.Train.Create))
	trainGroup.DELETE("/:name", util.Run(controller.Train.Delete))
	trainGroup.GET("/status/:name", util.Run(controller.Train.GetSingleStatus))
	trainGroup.POST("/status/list", util.Run(controller.Train.GetListStatus))
	trainGroup.GET("/log/:name", util.Run(controller.Train.GetLog))

	serviceGroup := e.Group("/api/service")
	serviceGroup.POST("", util.Run(controller.Service.Create))
	serviceGroup.DELETE("/:name", util.Run(controller.Service.Delete))
	serviceGroup.GET("/status/:name", util.Run(controller.Service.GetSingleStatus))
	serviceGroup.POST("/status/list", util.Run(controller.Service.GetListStatus))
	serviceGroup.GET("/log/:name", util.Run(controller.Service.GetLog))
}
