package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"runtime/debug"
)

type HandlerFunc func(*gin.Context) (interface{}, error)

type Result struct {
	Code  int         `json:"code"`
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}

func Run(handlerFunc HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if e := recover(); e != nil {
				debug.PrintStack()
				c.JSON(200, Result{
					Code:  500,
					Error: fmt.Sprintf("%s", e),
					Data:  nil,
				})
			}
		}()
		result, err := handlerFunc(c)
		if err != nil {
			err = errors.WithStack(err)
			fmt.Printf("%+v", err) // top two frame
			c.JSON(http.StatusOK, Result{
				Code:  400,
				Error: err.Error(),
				Data:  nil,
			})
			return
		}
		c.JSON(http.StatusOK, Result{
			Code:  0,
			Error: "",
			Data:  result,
		})
		return
	}
}

func RunStatus(handlerFunc HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if e := recover(); e != nil {
				debug.PrintStack()
				c.JSON(http.StatusForbidden, gin.H{
					"code":  500,
					"error": fmt.Sprintf("%s", e),
					"data":  nil,
				})
			}
		}()
		result, err := handlerFunc(c)
		if err != nil {
			err = errors.WithStack(err)
			fmt.Printf("%+v", err)
			c.JSON(http.StatusForbidden, gin.H{
				"code":  400,
				"error": err.Error(),
				"data":  nil,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code":  0,
			"error": nil,
			"data":  result,
		})
		return
	}
}

func Recovery(c *gin.Context, err interface{}) {
	debug.PrintStack()
	fmt.Printf("recover_err err:%+v\n", err)
	c.JSON(200, gin.H{
		"code":  500,
		"error": fmt.Sprintf("%s", err),
		"data":  nil,
	})
}
