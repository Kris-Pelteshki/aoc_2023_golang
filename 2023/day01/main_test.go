package main

import (
	"fmt"
	"testing"
)

var example1 = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

var example = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example1",
			input: example1,
			want:  142,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  281,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := part2(tt.input)
			if got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
			fmt.Println("Output:", got)
		})
	}
}
