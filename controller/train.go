package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"log"
	"model-k8s-service-go/logic"
)

type train struct {
}

var Train train

func (*train) Create(c *gin.Context) (interface{}, error) {
	var dto logic.TrainDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		log.Println("Bind请求参数失败, " + err.Error())
		return nil, errors.New("Bind请求参数失败")
	}
	svc, err := logic.Logic.CreateTrain(dto)
	if err != nil {
		return nil, err
	}
	return svc, nil
}

func (*train) Delete(c *gin.Context) (interface{}, error) {
	return logic.Logic.DeleteTrain(c.Param("name"))
}

func (*train) GetSingleStatus(c *gin.Context) (interface{}, error) {
	return logic.Logic.GetSingleStatus(c.Param("name"))
}

func (*train) GetLog(c *gin.Context) (interface{}, error) {
	return logic.Logic.GetLogs(c.Param("name"), c.DefaultQuery("line", "2000"))
}

func (*train) GetListStatus(c *gin.Context) (interface{}, error) {
	var list []string
	if err := c.ShouldBindJSON(&list); err != nil {
		log.Println("Bind请求参数失败, " + err.Error())
		return nil, errors.New("Bind请求参数失败")
	}
	return logic.Logic.GetListStatus(list)
}
