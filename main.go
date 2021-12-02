package main

import (
	"dbmock/db"
	parse "dbmock/flag"
	"dbmock/util"
	"fmt"
	"strings"
	"sync"
)

func main() {
	parse.Init()
	fmt.Println(parse.GetFilePath(),
		parse.GetPassord(),
		parse.GetUsername())

	util.Init()
	db.Init()
	do()
}

func do() {
	info := util.GetFileInfo()
	if info.MaxCount == 0 {
		fmt.Println("最大值不能为0,最小为50")
		return
	}
	if info.TableName == "" {
		fmt.Println("表名不能为空，且格式必须为 数据库.表名")
		return
	}
	if len(strings.Split(info.TableName, ".")) != 2 {
		fmt.Println("表名不能为空，且格式必须为 数据库.表名")
		return
	}
	if info.Cols == nil {
		fmt.Println("未指定数据库列名和数据类型")
		return
	}
	if info.Host == "" || info.Port == "" {
		fmt.Println("数据库链接配置不能为空")
		return
	}

	count := new(util.Counter)
	var wg sync.WaitGroup
	ch := make(chan struct{}, info.Advanced.OnceForGo)
	for i := 0; i <= info.MaxCount/info.Advanced.OnceForInsert; i++ {
		if count.Count() >= uint64(info.MaxCount) {
			break
		}
		ch <- struct{}{}
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			defer func() {
				<-ch
			}()
			if count.Count() >= uint64(info.MaxCount) {
				return
			}
			count.Incr(uint64(info.Advanced.OnceForInsert))
			err := db.Insert()
			if err != nil {
				var count int64
				_ = db.Count(info.TableName, &count)
				fmt.Println(fmt.Sprintf("插入失败，当前共插入 %d 条, 错误信息: %v", count, err))
				return
			}
		}(i)
	}
	wg.Wait()
}
