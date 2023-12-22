package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"log"
	"model-k8s-service-go/logic"
)

type service struct {
}

var Service service

func (*service) Create(c *gin.Context) (interface{}, error) {
	var dto logic.ServiceDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		log.Println("Bind请求参数失败, " + err.Error())
		return nil, errors.New("Bind请求参数失败")
	}

	svc, err := logic.Logic.CreateService(dto)
	if err != nil {
		return nil, err
	}
	return svc, nil
}

func (*service) Delete(c *gin.Context) (interface{}, error) {
	return logic.Logic.DeleteService(c.Param("name"))
}

func (*service) GetSingleStatus(c *gin.Context) (interface{}, error) {
	return logic.Logic.GetSingleStatus(c.Param("name"))
}

func (*service) GetListStatus(c *gin.Context) (interface{}, error) {
	var list []string
	if err := c.ShouldBindJSON(&list); err != nil {
		log.Println("Bind请求参数失败, " + err.Error())
		return nil, errors.New("Bind请求参数失败")
	}
	return logic.Logic.GetListStatus(list)
}

func (*service) GetLog(c *gin.Context) (interface{}, error) {
	return logic.Logic.GetLogs(c.Param("name"), c.DefaultQuery("line", "2000"))
}
