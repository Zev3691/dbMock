// Package db 数据库操作
package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	parse "dbmock/flag"
	"dbmock/util"
)

var db *gorm.DB

func Init() {
	NewDB(util.GetFileInfo())
}

func GetDB() *gorm.DB {
	return db
}

func NewDB(scan *util.Scan) {
	dsnTmpl := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(dsnTmpl, parse.GetUsername(), parse.GetPassord(),
		scan.Host, scan.Port, scan.Dbname)
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
}
