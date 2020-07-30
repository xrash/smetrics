package tests

import (
	"fmt"
	"github.com/xrash/smetrics"
	"testing"
)

func TestJaro(t *testing.T) {
	for _, c := range __jaro_cases {
		r := smetrics.Jaro(c.a, c.b)
		result := fmt.Sprintf("%.3f", r)
		expected := fmt.Sprintf("%.3f", c.r)
		if result != expected {
			fmt.Println(c.a, c.b, result, "instead of", expected)
			t.Fail()
		}
	}
}
