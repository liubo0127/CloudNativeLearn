package main

import "testing"

func Test_notFound1(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"test1", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := notFound1(); (err != nil) != tt.wantErr {
				t.Errorf("notFound1() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_notFound2(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"test2", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := notFound2(); (err != nil) != tt.wantErr {
				t.Errorf("notFound2() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
