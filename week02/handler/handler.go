package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type SingleHandler struct {
}

func (self *SingleHandler) ServeHTTP(http.ResponseWriter, *http.Request) {
	panic("implement me")
}

func NewSingleHandler() *SingleHandler {
	return &SingleHandler{}
}

func GetUsersHandler(c *gin.Context) {

}
