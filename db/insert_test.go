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
	scan := util.FileRead("C:\\Users\\99424\\project\\dbMock\\mock.json")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parseSQLForTest(scan)
			n := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(2)
			fmt.Println(n)
		})
	}
}

func BenchmarkParseSQL(b *testing.B) {
	scan := util.FileRead("C:\\Users\\99424\\project\\dbMock\\mock.json")
	for i := 0; i < b.N; i++ {
		parseSQLForTest(scan)
	}
}
