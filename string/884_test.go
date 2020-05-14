package string

import "testing"

func Test_myAtoi(t *testing.T) {
	tests := []struct {
		str  string
		want int
	}{
		{
			"42",
			42,
		},
		{
			"   -42",
			-42,
		},
		{
			" +-42",
			0,
		},
		{
			"4193 with words",
			4193,
		},
		{
			"words and 987",
			0,
		},
		{
			"-91283472332",
			-2147483648,
		},
		{
			"-2147483647",
			-2147483647,
		},
		{
			"-2147483648",
			-2147483648,
		},
		{
			"-2147483649",
			-2147483648,
		},
		{
			"2147483647",
			2147483647,
		},
		{
			"2147483648",
			2147483647,
		},
		{
			"981293812938129",
			2147483647,
		},
		{
			"     -981293812938129",
			-2147483648,
		},
	}
	for _, tt := range tests {
		t.Run(tt.str, func(t *testing.T) {
			if got := myAtoi(tt.str); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
