package dao

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"

	"github.com/intrsokx/geek-university/week02/config"
	"github.com/intrsokx/geek-university/week02/entity"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

var msql *gorm.DB

func InitMysql(dataSourceName string) {
	db, err := gorm.Open(config.Cfg.DB.Name, dataSourceName)
	if err != nil {
		panic(err)
	}

	db.DB().SetMaxOpenConns(200)
	db.DB().SetMaxIdleConns(100)

	msql = db
	logrus.Debug("msql == nil:", msql == nil)
}

/**
	如果在dao层发生了错误，将错误封装后，向上一层抛出。
错误的大致区分，在dao层是比较明显的，在将错误抛出后，由上层代码结构去处理错误。
*/
func GetUsers() ([]entity.User, error) {
	rows, err := msql.Table(config.Cfg.DB.Table.User).Rows()
	if err != nil {

		return nil, errors.Wrap(err, "[pkg dao] get users err")
	}

	ret := make([]entity.User, 0)

	for rows.Next() {
		u := entity.User{}
		err = rows.Scan(&u.NO, &u.Name, &u.Hobby, &u.Gender)
		if err != nil {
			return nil, errors.Wrap(err, "[pkg dao] scan rows err")
		}
		ret = append(ret, u)
	}

	return ret, nil
}
