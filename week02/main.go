package main

import (
	"database/sql"
	"github.com/intrsokx/geek-university/week02/handler"
	"github.com/sirupsen/logrus"
	"net/http"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:             true,
		TimestampFormat:           "2006-01-02 15:04:00",
	})
	logrus.SetReportCaller(true)
}

func main() {
	logrus.Info("hello world")
	_ = sql.ErrNoRows

	hd := handler.NewSingleHandler()

	http.Handle("/get/users", hd)

	logrus.Error(http.ListenAndServe(":8080", nil))
}
