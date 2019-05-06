package orm

import (
	"github.com/jinzhu/gorm"

	"github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MysqlConfig struct {
	User     string
	Password string
	Addr     string
	DBName   string
}

var mysqlInstance *gorm.DB

func Init(conf *MysqlConfig) {
	var (
		err error
		mc  mysql.Config
	)
	mc = mysql.Config{
		User:                 conf.User,
		Passwd:               conf.Password,
		Net:                  "tcp",
		Addr:                 conf.Addr,
		DBName:               conf.DBName,
		AllowNativePasswords: true,
	}
	if mysqlInstance, err = gorm.Open("mysql", mc.FormatDSN()); err != nil {
		panic("Mysql init error" + err.Error())
	}
}

func GetDB() *gorm.DB {
	return mysqlInstance
}
