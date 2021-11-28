package util

import "fmt"

var genData map[string]func(size int) string

func regiest() {
	genData = make(map[string]func(size int) string)
	genData["varchar"] = func(size int) string {
		return fmt.Sprintf("'%s'", genRandStr(size))
	}
	genData["char"] = func(size int) string {
		return fmt.Sprintf("'%s'", genRandStr(size))
	}
	genData["int"] = func(size int) string {
		if size >= 13 {
			return fmt.Sprintf("%d", genRandTime().UnixNano())
		} else {
			return fmt.Sprintf("%d", genRandTime().Unix())
		}
	}
	genData["tinyint"] = func(size int) string {
		return fmt.Sprintf("%d", seededRand.Intn(size))
	}
}

func GetData(typ string, size ...int) string {
	length := 0
	if len(size) > 0 {
		length = size[0]
	}
	return genData[typ](length)
}

func GetStrData(size int) string {
	return genData["varchar"](size)
}

func GetIntData(size int) string {
	return genData["int"](size)
}
