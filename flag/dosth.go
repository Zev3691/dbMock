package parse

import "flag"

type path struct{}

func (r *path) Parse() {
	flag.StringVar(&filepath, "path", "", "解析文件的路径")
}

type dbAuth struct{}

func (r *dbAuth) Parse() {
	flag.StringVar(&username, "u", "", "数据库用户名")
	flag.StringVar(&password, "p", "", "数据库密码")
	flag.IntVar(&isEncode, "e", 0, "数据库密码加密类型，0/1 支持base64")
}

func init() {
	regiestParse(&path{})
	regiestParse(&dbAuth{})
}
