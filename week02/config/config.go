package config

import (
	"github.com/jinzhu/configor"
	"github.com/sirupsen/logrus"
)

var Cfg = struct {
	Server struct {
		Name string
		Port int
	}

	DB struct {
		Name string
		User string `default:"root"`
		//Password string `required:"true" env:"DBPassword"`
		Password string
		Database string

		Table struct {
			User string
		}
	}
}{}

func InitConf(name ...string) {
	err := configor.Load(&Cfg, name...)
	if err != nil {
		panic(err)
	}
	logrus.Debugf("%+v", Cfg)
}
