package db

import (
	"dbmock/util"
	"fmt"
	"strings"
)

func parseSQL(scan *util.Scan) string {
	return ""
}

func parseSQLForTest(scan *util.Scan) string {
	// sqlTmpl := "INSERT INTO %s SET %s VALUES %s"
	vals := []string{}
	col := ""
	// info := util.GetFileInfo()
	for k, v := range scan.Cols {
		col += k + ","
		vals = append(vals, v)
	}
	col = strings.Trim(col, ",")
	fmt.Println(col)
	fmt.Println(vals)
	valStr := ""
	for _, v := range vals {
		valStr += util.GetData(v)
	}
	fmt.Println(valStr)
	// sql := fmt.Sprintf(sqlTmpl, info.TableName)
	return ""
}
