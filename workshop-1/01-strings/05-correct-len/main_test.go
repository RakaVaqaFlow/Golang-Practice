package main

import "testing"

func TestIsStringShorterOrEqualThen3(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsStringShorterOrEqualThen3(tt.args.s); got != tt.want {
				t.Errorf("IsStringShorterOrEqualThen3() = %v, want %v", got, tt.want)
			}
		})
	}
}
