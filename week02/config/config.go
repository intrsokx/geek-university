package config

import (
	"github.com/jinzhu/configor"
	"github.com/sirupsen/logrus"
)

var Config = struct {
	Server struct{
		Name string `default:"app name"`
		Port int `default:"8097"`
	}

	DB struct {
		Name     string
		User     string `default:"root"`
		Password string `required:"true" env:"DBPassword"`
		Port     uint   `default:"3306"`
	}
}{}

func InitConf(name ...string)  {
	configor.Load(&Config, name...)
	logrus.Debugf("%+v", Config)
}