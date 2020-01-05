package main

import "testing"

func Test_cashback(t *testing.T) {

	tests := []struct {
		name string
		amount int
		want int
	}{
		// TODO: Add test cases.
		{"Have cashback", 5000, 250},
		{"No cashback", 1000, 0},
		{"Cshback", 3000, 150},

	}
	for _, test := range tests {
		got:= cashback(test.amount)
		if got != test.want {
			t.Error("For cashback test", test.name, "got: ", got, "want: ", test.want)
		}
	}
}