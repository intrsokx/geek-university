package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/intrsokx/geek-university/week02/dao"
	"github.com/sirupsen/logrus"
	"net/http"
)

/**
1、在handler层处理具体的业务逻辑
2、在这里，我打印了具体的调用链日志信息。并且将不带调用栈信息的错误返回给前端
*/
func GetUsersHandler(c *gin.Context) {
	users, err := dao.GetUsers()
	if err != nil {
		//%+v 打印错误信息及堆栈调用信息
		logrus.Errorf("%+v", err)
		c.String(http.StatusOK, "get user err:%v\n", err)
		return
	}
	c.JSON(http.StatusOK, users)
}
