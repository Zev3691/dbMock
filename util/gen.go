package util

import (
	"fmt"
	"math/rand"
	"time"
)

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func init() {
	regiest()
}

// genRandStr 生成指定长度的随机字符串
func genRandStr(length int) string {
	if length == 0 {
		length = 10
	}
	return stringWithCharset(length, retCharSet())
}

// genRandTime 生成随机的时间
// 数字16为系数.可随意调整以达到目的，目的：防止生产的时间戳太大，
// 建议设置倍数：2，4，8，16
// 也可计算：rand.Int63n(time.Now().Unix()-94608000) + 94608000
func genRandTime() time.Time {
	randomTime := rand.Int63n(time.Now().Unix()/16) + time.Now().Unix()
	randomNow := time.Unix(randomTime, 0)
	return randomNow
}

func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		l := len(charset)
		fmt.Println(l)
		b[i] = charset[seededRand.Intn(l)]
	}
	return string(b)
}

func retCharSet() string {
	return "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
}
