package tests

import (
	"fmt"
	"github.com/xrash/smetrics"
	"testing"
)

func TestUkkonen(t *testing.T) {
	cases := []levenshteincase{
		{"", "abc", 0, 1, 2, 0},
		{"abc", "", 1, 0, 2, 0},
		{"abc", "xyz", 1, 1, 0, 0},
		{"ab", "cab", 0, 2, 3, 0},
		{"cab", "ab", 2, 0, 3, 0},
		{"ab", "xyz", 2, 3, 0, 2},
		{"abc", "xyz", 0, 0, 0, 0},
		{"", "", 0, 1, 1, 0},
		{"abc", "xyz", -1, 1, 2, 0},
		{"abc", "xyz", 1, -1, 2, 0},
		{"abc", "xyz", 1, 1, -1, -3},
		{"ab", "xyz", -1, -2, -3, -7},
		{"", "abc", -1, 1, 1, -3},
		{"abc", "def", -1, 2, 3, 3},
		{"abcd", "x", 1, -1, 2, -3},
		{"", "", -1, 0, 1, 0},
		{"RASH", "RASH", 1, 1, 2, 0},
		{"POTATO", "POTTATO", 1, 1, 2, 1},
		{"POTTATO", "POTATO", 1, 1, 2, 1},
		{"HOUSE", "MOUSE", 1, 1, 2, 2},
		{"MOUSE", "HOUSE", 2, 2, 4, 4},
		{"abc", "xy", 2, 3, 5, 13},
		{"xy", "abc", 2, 3, 5, 12},
	}

	for _, c := range cases {
		if r := smetrics.Ukkonen(c.s, c.t, c.icost, c.dcost, c.scost); r != c.r {
			fmt.Println(r, "instead of", c.r)
			t.Fail()
		}
	}
}
