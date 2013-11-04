package tests

import (
	"testing"
	"fmt"
	"github.com/xrash/smetrics"
)

func TestUkkonen(t *testing.T) {
	cases := []levenshteincase{
		{"RASH", "RASH", 1, 1, 2, 0},
		{"POTATO", "POTTATO", 1, 1, 2, 1},
		{"POTTATO", "POTATO", 1, 1, 2, 1},
		{"HOUSE", "MOUSE", 1, 1, 2, 2},
		{"MOUSE", "HOUSE", 2, 2, 4, 4},
	}

	for _, c := range cases {
		if r := smetrics.Ukkonen(c.s, c.t, c.icost, c.dcost, c.scost); r != c.r {
			fmt.Println(r, "instead of", c.r)
			t.Fail()
		}
	}
}
