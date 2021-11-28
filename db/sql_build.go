package db

import (
	"dbmock/util"
	"fmt"
	"strings"
)

func parseSQL(scan *util.Scan) string {
	sqlTmpl := "INSERT INTO %s %s VALUES %s"
	var vals []string
	col := "("
	for k, v := range scan.Cols {
		col += k + ","
		vals = append(vals, v)
	}
	col = strings.TrimSuffix(col, ",") + ")"
	var valStr []string
	for i := 0; i <= scan.Advanced.OnceForInsert; i++ {
		val := "("
		for _, v := range vals {
			val += util.GetData(v)+","
		}
		val = strings.TrimSuffix(val, ",") + ")"
		valStr = append(valStr, val)
	}
	return fmt.Sprintf(sqlTmpl, scan.TableName, col, strings.Join(valStr, ","))
}

// 测试使用
func parseSQLForTest(scan *util.Scan) string {
	sqlTmpl := "INSERT INTO %s SET %s VALUES %s"
	var vals []string
	col := "("
	for k, v := range scan.Cols {
		col += k + ","
		vals = append(vals, v)
	}
	col = strings.TrimSuffix(col, ",") + ")"
	var valStr []string
	for i := 0; i <= scan.Advanced.OnceForInsert; i++ {
		val := "("
		for _, v := range vals {
			val += util.GetData(v)+","
		}
		val = strings.TrimSuffix(val, ",") + ")"
		valStr = append(valStr, val)
	}
	fmt.Println(valStr)
	sql := fmt.Sprintf(sqlTmpl, scan.TableName, col, strings.Join(valStr, ","))
	fmt.Println(sql)
	return ""
}
