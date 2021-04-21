package main

import (
	"github.com/gin-gonic/gin"
	"github.com/intrsokx/geek-university/week02/config"
	"github.com/intrsokx/geek-university/week02/handler"
	"github.com/sirupsen/logrus"
	"net/http"
)

func initLog() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:             true,
		TimestampFormat:           "2006-01-02 15:04:00",
	})
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.DebugLevel)
}

func init() {
	initLog()

	config.InitConf("config.yml")
}

func main() {
	e := gin.New()
	e.Use(gin.Logger())
	e.Use(gin.Recovery())


	e.Handle(http.MethodGet, "/get/users", handler.GetUsersHandler)

	logrus.Error(e.Run())
}
