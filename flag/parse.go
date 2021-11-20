// Package parse 实现了命令行参数的注册
package parse

import "flag"

type parse interface {
	Parse()
}

var parseFunc []parse

// 调用使用init，确保在包的Init函数之前注册完毕
func regiestParse(f parse) {
	if parseFunc == nil {
		parseFunc = make([]parse, 0)
	}
	parseFunc = append(parseFunc, f)
}

func Init() {
	for _, f := range parseFunc {
		f.Parse()
	}
	flag.Parse()
}
