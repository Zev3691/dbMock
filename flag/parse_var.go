package parse

// 定义命令行解析参数，并设置getter

var (
	filepath string
	username string
	password string
	isEncode int
)

func GetFilePath() string {
	return filepath
}

func GetUsername() string {
	return username
}

func GetPassord() string {
	return password
}

func GetIsEncode() int {
	return isEncode
}
