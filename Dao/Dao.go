package Dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

var DB *gorm.DB
var err error

/**
链接数据库
 */
func InitMysql() {
	dsn:= "root:uChB;r6rdmo-!@tcp(mine:3306)/django?charset=utf8&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql",dsn)
	DB.LogMode(true)
	if err != nil {
		fmt.Println("db链接失败", err.Error())
		os.Exit(1)
	} else {
		fmt.Println("db connect success")
	}
}
/**
关闭数据库
 */
func Close() {
	DB.Close()
}
