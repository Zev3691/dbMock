// Package util 公共函数
package util

import parse "dbmock/flag"

var scan *Scan

func Init() {
	scan = fileUnmarshal(parse.GetFilePath())
}

func GetFileInfo() *Scan {
	return scan
}
