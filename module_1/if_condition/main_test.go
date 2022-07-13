package main

import "testing"

func Test_ifAssignValue(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test1", args{11}, 10},
		{"test2", args{6}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ifAssignValue(tt.args.val); got != tt.want {
				t.Errorf("ifAssignValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ifStandard(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test1", args{11}, 10},
		{"test2", args{6}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ifStandard(tt.args.val); got != tt.want {
				t.Errorf("ifStandard() = %v, want %v", got, tt.want)
			}
		})
	}
}
