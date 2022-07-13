package main

import "testing"

func Test_switchStandard(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"test1", args{1}},
		{"test2", args{2}},
		{"test3", args{3}},
		{"test4", args{4}},
		{"test5", args{5}},
		{"test6", args{6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
