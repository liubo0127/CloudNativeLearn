package main

import "testing"

func Test_priAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test1", args{1, 2}, 3},
		{"test2", args{3,5},8},
		{"test3", args{2,3}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := priAdd(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("priAdd() = %v, want %v", got, tt.want)
			}
		})
	}
}
