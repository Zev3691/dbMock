package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

type Scan struct {
	TableName string            `json:"table_name,omitempty"`
	Cols      map[string]string `json:"cols,omitempty"`
	MaxCount  int               `json:"max_count,omitempty"`
	Host      string            `json:"host,omitempty"`
	Port      string            `json:"port,omitempty"`
	Dbname    string            `json:"dbname,omitempty"`
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
	j.readFile()
	var scan Scan
	json.Unmarshal(j.body, &scan)
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
