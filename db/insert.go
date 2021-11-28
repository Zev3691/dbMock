package db

import "dbmock/util"

func Insert() error {
	db := GetDB()
	sql := parseSQL(util.GetFileInfo())
	return db.Exec(sql).Error
}


func Count(tableName string, count *int64) error {
	db := GetDB()
	return db.Table(tableName).Count(count).Error
}