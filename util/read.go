package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

type Scan struct {
	TableName string            `json:"table_name"`
	Cols      map[string]string `json:"cols"`
	MaxCount  int               `json:"max_count"`
	Host      string            `json:"host"`
	Port      string            `json:"port"`
	Dbname    string            `json:"dbname"`
	Advanced  advanced          `json:"advanced"`
}

// 看到这里的同学可以在mock.json中的增加相关结构,
// once_for_go 默认100
// once_for_insert 默认50
// 需要注意数据库的性能是否匹配下面的配置
type advanced struct {
	OnceForGo     int `json:"once_for_go"`
	OnceForInsert int `json:"once_for_insert"`
}

type Read interface {
	Unmarshal()
	Marshal()
}

type JSON struct {
	Path string // 外部输入，使用绝对路径效果更佳
	body []byte
}

func (j *JSON) readFile() error {
	if j.Path == "" {
		return errors.New("未指定配置文件目录！")
	}
	fb, err := ioutil.ReadFile(j.Path)
	if err != nil {
		return fmt.Errorf("读取配置文件错误： %v", err.Error())
	}
	j.body = fb
	return nil
}

func (j *JSON) Marshal() {
	// 预留 什么都不做
}

func (j *JSON) Unmarshal() *Scan {
	_ = j.readFile()
	var scan Scan
	_ = json.Unmarshal(j.body, &scan)
	if scan.Advanced.OnceForInsert == 0 {
		scan.Advanced.OnceForInsert = 50
	}
	if scan.Advanced.OnceForGo == 0 {
		scan.Advanced.OnceForGo = 100
	}
	return &scan
}

func fileUnmarshal(path string) *Scan {
	pathSlc := strings.Split(path, "/")
	if len(pathSlc) == 0 {
		fmt.Printf("不能识别的路径\n")
		return nil
	}
	fileType := strings.Split(pathSlc[len(pathSlc)-1], ".")
	if len(fileType) == 0 || len(fileType) == 1 {
		return nil
	}

	switch fileType[1] {
	case "json":
		model := &JSON{Path: path}
		return model.Unmarshal()
	}
	return nil
}

func FileRead(path string) *Scan {
	return fileUnmarshal(path)
}
