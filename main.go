package main

import (
	"encoding/json"
	"fmt"
)

type Scan struct {
	TableName string            `json:"table_name,omitempty"`
	Cols      map[string]string `json:"cols,omitempty"`
	MaxCount  int               `json:"max_count,omitempty"`
}

func main() {
	jsStr := `
{
	"table_name":"list",
	"cols": {
		"id":"int",
		"name":"varchar",
		"addr":"varchar"
	},
	"max_count": 10
}
	`
	var js Scan
	json.Unmarshal([]byte(jsStr), &js)
	fmt.Println(js)
}
