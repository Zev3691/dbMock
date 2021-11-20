package db

import (
	"dbmock/util"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestParseSQL(t *testing.T) {
	// 参数
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"round 1"},
		// {"round 2"},
		// {"round 3"},
		// {"round 4"},
	}
	scan := util.FileRead("/home/whh/GolandProjects/dbMock/mock.json")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parseSQLForTest(scan)
			n := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(2)
			fmt.Println(n)
		})
	}
}
