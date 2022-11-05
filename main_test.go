package main

import (
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_function_which_return_1(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "test1",
			want: 1,
		},
		{
			name: "test2",
			want: 1,
		},
		{
			name: "test3",
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := function_which_return_1(); got != tt.want {
				t.Errorf("function_which_return_1() = %v, want %v", got, tt.want)
			}
		})
	}
}
