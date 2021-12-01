package util

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestGenRandInt(t *testing.T) {
	// 参数
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"size1", args{4}},
		{"size2", args{5}},
		{"size3", args{6}},
		{"size4", args{7}},
	}
	for _, tt := range tests {
		var seededRand = rand.New(
			rand.NewSource(time.Now().UnixNano()))
		t.Run(tt.name, func(t *testing.T) {
			fmt.Printf("%s: %d\n", tt.name, seededRand.Intn(9))
		})
	}
}

func Test_genRandTime(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"1"},
		{"1"},
		{"1"},
		{"1"},
		{"1"},
		{"1"},
		{"1"},
		{"1"},
		{"1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := genRandTime()
			fmt.Println(tt.Unix())
		})
	}
}
