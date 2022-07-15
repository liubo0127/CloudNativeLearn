package main

import "testing"

func Test_personInfo_getAddress(t *testing.T) {
	type fields struct {
		age     string
		address string
	}
	tests := []struct {
		name        string
		fields      fields
		wantAddress string
	}{
		// TODO: Add test cases.
		{"zhangsan", fields{"18", "北京"}, "北京"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &personInfo{
				age:     tt.fields.age,
				address: tt.fields.address,
			}
			if gotAddress := p.getAddress(); gotAddress != tt.wantAddress {
				t.Errorf("getAddress() = %v, want %v", gotAddress, tt.wantAddress)
			}
		})
	}
}

func Test_personInfo_getAge(t *testing.T) {
	type fields struct {
		age     string
		address string
	}
	tests := []struct {
		name    string
		fields  fields
		wantAge string
	}{
		// TODO: Add test cases.
		{"test1", fields{"18", "北京"}, "18"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &personInfo{
				age:     tt.fields.age,
				address: tt.fields.address,
			}
			if gotAge := p.getAge(); gotAge != tt.wantAge {
				t.Errorf("getAge() = %v, want %v", gotAge, tt.wantAge)
			}
		})
	}
}
